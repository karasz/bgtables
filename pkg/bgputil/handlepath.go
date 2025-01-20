// Package bgputil contains utilities to handle BGP features
package bgputil

import (
	"fmt"
	"net"

	apipb "github.com/osrg/gobgp/v3/api"
	"github.com/osrg/gobgp/v3/pkg/apiutil"
	"github.com/osrg/gobgp/v3/pkg/packet/bgp"

	"github.com/karasz/bgtables/pkg/kernel"
)

const (
	afiIP               = apipb.Family_AFI_IP
	afiL2VPN            = apipb.Family_AFI_L2VPN
	safiUnicast         = apipb.Family_SAFI_UNICAST
	safiEVPN            = apipb.Family_SAFI_EVPN
	safiFlowSpecUnicast = apipb.Family_SAFI_FLOW_SPEC_UNICAST
)

type unsupportedAFISAFIError struct {
	afi  apipb.Family_Afi
	safi apipb.Family_Safi
}

// Error returns a formatted error message indicating the unsupported AFI/SAFI values.
// The message includes the AFI (Address Family Identifier) and SAFI (Subsequent Address Family Identifier) values.
func (e *unsupportedAFISAFIError) Error() string {
	return fmt.Sprintf("unsupported AFI/SAFI: %d/%d", e.afi, e.safi)
}

// HandlePath processes a BGP path based on its AFI/SAFI values and the specified neighbor.
// It supports handling for IP Unicast, EVPN, and Flow Spec Unicast paths.
// Returns an error if the AFI/SAFI is unsupported or if any processing step fails.
//
// Parameters:
// - path: A pointer to an apipb.Path object representing the BGP path to be processed.
//
// Returns:
// - error: An error object if any processing step fails or if the AFI/SAFI is unsupported, otherwise nil.
func HandlePath(path *apipb.Path) error {
	// Check if path is nil
	if path == nil {
		return fmt.Errorf("path is nil")
	}
	// Validate path fields
	if path.GetFamily() == nil {
		return fmt.Errorf("path family is nil")
	}

	// Extract AFI/SAFI from the path
	pafi := path.GetFamily().GetAfi()   // AFI (e.g., IPv4, IPv6)
	psafi := path.GetFamily().GetSafi() // SAFI (e.g., Unicast, MPLS VPN)

	if pafi == 0 || psafi == 0 {
		return fmt.Errorf("AFI or SAFI is not set")
	}
	// Map of handlers for AFI/SAFI pairs
	handlers := map[struct {
		afi  apipb.Family_Afi
		safi apipb.Family_Safi
	}]func(*apipb.Path) error{
		{afiIP, safiUnicast}:         handleIPAddrPrefix,
		{afiL2VPN, safiEVPN}:         handleEVPN,
		{afiIP, safiFlowSpecUnicast}: handleFlowSpec,
	}

	// Find and execute the handler
	if handler, found := handlers[struct {
		afi  apipb.Family_Afi
		safi apipb.Family_Safi
	}{pafi, psafi}]; found {
		return handler(path)
	}
	return &unsupportedAFISAFIError{afi: pafi, safi: psafi}
}

// handleIPAddrPrefix processes a BGP path with IP address prefixes.
func handleIPAddrPrefix(path *apipb.Path) error {
	if err := validatePath(path); err != nil {
		return err
	}

	nlri, err := getNLRI(path)
	if err != nil {
		return err
	}

	nextHop, dstSubnet, err := getNexthop(path)
	if err != nil {
		return err
	}

	return updateKernelRoute(nlri, nextHop, dstSubnet)
}

// validatePath checks if the path and its family are valid.
func validatePath(path *apipb.Path) error {
	if path == nil {
		return fmt.Errorf("path is nil")
	}
	if path.GetFamily() == nil {
		return fmt.Errorf("path family is nil")
	}
	return nil
}

// getNLRI retrieves the NLRI from the path.
func getNLRI(path *apipb.Path) (bgp.AddrPrefixInterface, error) {
	ntu := path.GetNlri().TypeUrl
	if ntu != "type.googleapis.com/gobgpapi.IPAddressPrefix" {
		return nil, fmt.Errorf("invalid type URL: %s", ntu)
	}

	nlri, err := apiutil.GetNativeNlri(path)
	if err != nil {
		return nil, fmt.Errorf("failed to get native NLRI: %v", err)
	}

	return nlri, nil
}

// updateKernelRoute updates the kernel routing table with the given NLRI, next hop, and destination subnet.
func updateKernelRoute(nlri bgp.AddrPrefixInterface, nextHop net.IP, dstSubnet []net.IPNet) error {
	if nextHop == nil {
		return fmt.Errorf("failed to parse next hop from path NLRI: %v", nlri)
	}

	dstS, family := parseNLRI(nlri)
	if (dstS.IP != nil && dstS.Mask != nil) && dstSubnet != nil {
		dstSubnet = append(dstSubnet, dstS)
	}

	if dstSubnet != nil {
		kernel.UpdateKernelRoute(family, dstSubnet, nextHop)
	}

	return nil
}

// getNexthop extracts the next hop IP address from the given BGP path attributes.
func getNexthop(path *apipb.Path) (net.IP, []net.IPNet, error) {
	attrs := path.GetPattrs()

	for _, at := range attrs {
		n, err := apiutil.UnmarshalAttribute(at)
		if err != nil {
			return nil, nil, err
		}

		nextHop, prefixes, err := extractNextHopAndPrefixes(n)
		if err == nil {
			return nextHop, prefixes, nil
		} else if err != errNoNextHop {
			return nil, nil, err
		}
	}

	return nil, nil, fmt.Errorf("no next hop attribute found")
}

var errNoNextHop = fmt.Errorf("no next hop available")

// extractNextHopAndPrefixes processes a BGP path attribute to extract next hop and prefixes.
func extractNextHopAndPrefixes(attr bgp.PathAttributeInterface) (net.IP, []net.IPNet, error) {
	switch v := attr.(type) {
	case *bgp.PathAttributeNextHop:
		return v.Value, nil, nil
	case *bgp.PathAttributeMpReachNLRI:
		return processMpReachNLRI(v)
	default:
		return nil, nil, errNoNextHop
	}
}

// processMpReachNLRI processes a PathAttributeMpReachNLRI to extract next hop and prefixes.
func processMpReachNLRI(attr *bgp.PathAttributeMpReachNLRI) (net.IP, []net.IPNet, error) {
	prefixes, err := extractPrefixes(attr)
	if err != nil {
		return nil, nil, err
	}
	if len(attr.Nexthop) > 0 {
		return net.IP(attr.Nexthop), prefixes, nil
	}
	return nil, nil, fmt.Errorf("MPReachNLRI attribute found but no next hop available")
}

// extractPrefixes extracts IP prefixes from a PathAttributeMpReachNLRI attribute.
func extractPrefixes(attr *bgp.PathAttributeMpReachNLRI) ([]net.IPNet, error) {
	var prefixes []net.IPNet

	for _, nlri := range attr.Value {
		prefix, err := convertNLRIToPrefix(nlri)
		if err != nil {
			return nil, err
		}
		prefixes = append(prefixes, prefix)
	}

	return prefixes, nil
}

// convertNLRIToPrefix converts an NLRI to a net.IPNet prefix.
func convertNLRIToPrefix(nlri bgp.AddrPrefixInterface) (net.IPNet, error) {
	switch p := nlri.(type) {
	case *bgp.IPAddrPrefix:
		ip := net.IP(p.Prefix)
		mask := net.CIDRMask(int(p.Length), 8*len(ip))
		return net.IPNet{IP: ip, Mask: mask}, nil
	case *bgp.LabeledIPAddrPrefix:
		ip := net.IP(p.Prefix)
		mask := net.CIDRMask(int(p.Length), 8*len(ip))
		return net.IPNet{IP: ip, Mask: mask}, nil
	default:
		return net.IPNet{}, fmt.Errorf("unsupported NLRI type: %T", nlri)
	}
}

func parseNLRI(nlri bgp.AddrPrefixInterface) (net.IPNet, string) {
	prefix := prefixInterface2netIP(nlri)
	switch nlri.(type) {
	case *bgp.IPAddrPrefix:
		return prefix, ""
	case *bgp.IPv6AddrPrefix:
		return prefix, "-6"
	}
	return prefix, ""
}

func prefixInterface2netIP(v bgp.AddrPrefixInterface) net.IPNet {
	var prefix net.IPNet
	switch z := v.(type) {
	case *bgp.IPAddrPrefix:
		ip := net.IP(z.Prefix)
		mask := net.CIDRMask(int(z.Length), 8*len(ip)) // Create the subnet mask
		prefix.IP = ip
		prefix.Mask = mask
	case *bgp.IPv6AddrPrefix:
		ip := net.IP(z.Prefix)
		mask := net.CIDRMask(int(z.Length), 8*len(ip)) // Create the subnet mask
		prefix.IP = ip
		prefix.Mask = mask
	}
	return prefix
}
func handleFlowSpec(path *apipb.Path) error {
	_, _ = fmt.Println(path)
	return nil
}
func handleEVPN(path *apipb.Path) error {
	/*
		nlri := path.GetNlri()
		if nlri == nil {
			return fmt.Errorf("failed to parse NLRI from path: %v", path)
		}

		// Type assertion to EVPN NLRI
		evpnNlri, ok := interface{}(nlri).(*bgp.EVPNNLRI)
		if !ok {
			return fmt.Errorf("not an EVPN NLRI")
		}

		// Switch based on EVPN route type
		switch route := evpnNlri.RouteTypeData.(type) {
		case *bgp.EVPNIPPrefixRoute:
			fmt.Printf("EVPN IP Prefix Route: %s\n", route.IPPrefix.String())
		case *bgp.EVPNMacIPAdvertisementRoute:
			fmt.Printf("EVPN MAC/IP Route: MAC=%s, IP=%s\n",
				route.MacAddress.String(), route.IPAddress.String())
		case *bgp.EVPNMulticastEthernetTagRoute:
			fmt.Printf("EVPN Multicast Route: EthernetTag=%d\n", route.ETag)
		case *bgp.EVPNEthernetAutoDiscoveryRoute:
			fmt.Printf("EVPN Auto-Discovery Route: RD=%s, EthernetTag=%d\n",
				route.RD.String(), route.ETag)
		default:
			return fmt.Errorf("unsupported EVPN route type: %T", route)
		}
		return nil
	*/
	_, _ = fmt.Println(path)
	return nil
}

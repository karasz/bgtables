// Package routes provides functions to fetch and update BGP routes.
package routes

import (
	"context"
	"fmt"
	"net"

	apipb "github.com/osrg/gobgp/v3/api"
	"google.golang.org/protobuf/types/known/anypb"
)

// FetchRoutes fetches all routes from the GoBGP server.
func FetchRoutes(client apipb.GobgpApiClient) ([]*apipb.Path, error) {
	stream, err := client.ListPath(context.Background(), &apipb.ListPathRequest{
		TableType: apipb.TableType_GLOBAL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list paths: %w", err)
	}

	var paths []*apipb.Path
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}

		paths = append(paths, resp.Destination.Paths...)
	}

	return paths, nil
}

// ParseNlriToCIDR decodes the NLRI from *anypb.Any to a string in CIDR format.
func ParseNlriToCIDR(nlri *anypb.Any) (string, error) {
	prefix := &apipb.IPAddressPrefix{}
	if err := nlri.UnmarshalTo(prefix); err != nil {
		return "", fmt.Errorf("failed to unmarshal NLRI: %w", err)
	}

	addr, err := net.ResolveIPAddr("ip", prefix.Prefix)
	if err != nil {
		return "", err
	}

	ipStr := addr.IP.String()

	cidr := fmt.Sprintf("%s/%d", ipStr, prefix.PrefixLen)

	_, _, err = net.ParseCIDR(cidr)

	if err != nil {
		return "", fmt.Errorf("invalid CIDR format: %w", err)
	}

	return cidr, nil
}

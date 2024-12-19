package routes

import (
	"fmt"
	"log"

	apipb "github.com/osrg/gobgp/v3/api"
	"github.com/vishvananda/netlink"
)

type routeOperation struct {
	cidr  string
	route *netlink.Route
}

// UpdateLocalRoutes updates the local routes with the provided paths.
func UpdateLocalRoutes(paths []*apipb.Path) error {
	existingRoutes, err := getExistingRoutes()
	if err != nil {
		return err
	}

	desiredRoutes := buildDesiredRoutes(paths)

	if err := applyRouteChanges(existingRoutes, desiredRoutes); err != nil {
		return fmt.Errorf("failed to apply route changes: %w", err)
	}

	return nil
}

func getExistingRoutes() (map[string]*netlink.Route, error) {
	routes, err := netlink.RouteList(nil, netlink.FAMILY_ALL)
	if err != nil {
		return nil, fmt.Errorf("failed to list kernel routes: %w", err)
	}

	routeMap := make(map[string]*netlink.Route)
	for _, route := range routes {
		if route.Dst != nil {
			routeMap[route.Dst.String()] = &route
		}
	}
	return routeMap, nil
}

func buildDesiredRoutes(paths []*apipb.Path) map[string]*netlink.Route {
	desiredRoutes := make(map[string]*netlink.Route)

	for _, path := range paths {
		if path.IsWithdraw {
			continue
		}

		route := createRouteFromPath(path)
		if route != nil {
			desiredRoutes[route.cidr] = route.route
		}
	}

	return desiredRoutes
}

func createRouteFromPath(path *apipb.Path) *routeOperation {
	cidr, err := ParseNlriToCIDR(path.Nlri)
	if err != nil {
		log.Printf("Failed to parse Nlri %v: %v", path.Nlri, err)
		return nil
	}

	dst, err := netlink.ParseIPNet(cidr)
	if err != nil {
		log.Printf("Invalid route CIDR %s for addition: %v", cidr, err)
		return nil
	}

	return &routeOperation{
		cidr:  cidr,
		route: &netlink.Route{Dst: dst},
	}
}

func applyRouteChanges(existing, desired map[string]*netlink.Route) error {
	if err := addOrUpdateRoutes(desired); err != nil {
		return err
	}

	return removeStaleRoutes(existing, desired)
}

func addOrUpdateRoutes(desired map[string]*netlink.Route) error {
	for cidr, route := range desired {
		if err := updateRoute(cidr, route); err != nil {
			log.Printf("Failed to manage route %s: %v", cidr, err)
		}
	}
	return nil
}

func removeStaleRoutes(existing, desired map[string]*netlink.Route) error {
	for cidr, route := range existing {
		if _, exists := desired[cidr]; !exists {
			if err := removeRoute(cidr, route); err != nil {
				log.Printf("Failed to remove route %s: %v", cidr, err)
			}
		}
	}
	return nil
}

func updateRoute(cidr string, route *netlink.Route) error {
	if err := netlink.RouteReplace(route); err != nil {
		return fmt.Errorf("failed to update route: %w", err)
	}
	log.Printf("Updated route: %s", cidr)
	return nil
}

func removeRoute(cidr string, route *netlink.Route) error {
	if err := netlink.RouteDel(route); err != nil {
		return fmt.Errorf("failed to delete route: %w", err)
	}
	log.Printf("Removed route: %s", cidr)
	return nil
}

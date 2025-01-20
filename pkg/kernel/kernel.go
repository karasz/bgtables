// Package kernel manipulates kernel routes on linux machines
package kernel

import (
	"log"
	"net"
	"os/exec"
)

// UpdateKernelRoute adds or updates a route in the kernel routing table.
// It executes the "ip route replace" command with the specified family, prefix, and nexthop.
// The "replace" command is used to add a new route or update an existing one.
func UpdateKernelRoute(family string, prefixes []net.IPNet, nexthop net.IP) {
	for _, prefix := range prefixes {
		cmd := exec.Command("ip", family, "route", "replace", prefix.String(), "via", nexthop.String())
		if err := cmd.Run(); err != nil {
			log.Printf("Failed to update kernel route: %v", err)
			continue
		}
		log.Printf("Updated kernel route: %s via %s", prefix, nexthop)
	}
}

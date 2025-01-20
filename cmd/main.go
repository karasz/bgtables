// Package main - Entry point for the BGP tables application.
//
// This file contains the main function which initializes and starts the BGP server,
// connects to BGP peers, and watches for BGP updates.
//
// Functions:
// - main: Loads the configuration, creates and starts the BGP server, connects to peers
// and starts watching for BGP updates.
// - watchBgpUpdates: Watches for BGP events in the best path table and handles path updates.
//
// Dependencies:
// - github.com/karasz/bgtables/config: For loading configuration.
// - github.com/karasz/bgtables/pkg/bgpserver: For creating and managing the BGP server.
// - github.com/karasz/bgtables/pkg/bgputil: For handling BGP path updates.
// - github.com/osrg/gobgp/v3/api: For BGP API definitions and requests.
package main

import (
	"context"
	"log"

	"github.com/karasz/bgtables/config"
	"github.com/karasz/bgtables/pkg/bgpserver"
	"github.com/karasz/bgtables/pkg/bgputil"
	apipb "github.com/osrg/gobgp/v3/api"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create a new BGP server instance
	bgpSrv, err := bgpserver.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create BGP server: %v", err)
	}

	go bgpSrv.Serve()
	// Start the BGP server
	// Magic ASN. First free after the kube-router ASN
	// 255 should work for routerid
	// Do not listen on any port, we are just "a client"
	if err := bgpSrv.StartBgp(context.Background(), &apipb.StartBgpRequest{
		Global: &apipb.Global{
			Asn:        64513,
			RouterId:   "255.255.255.255",
			ListenPort: -1,
		},
	}); err != nil {
		log.Fatalf("Failed to start BGP server: %v", err)
	}
	// Connect to the BGP peers
	for _, p := range cfg.Peers {
		peer := &apipb.Peer{
			Conf: &apipb.PeerConf{
				PeerAsn:         p.Asn,
				NeighborAddress: p.IP,
			},
		}
		if err := bgpSrv.AddPeer(context.Background(), &apipb.AddPeerRequest{
			Peer: peer,
		}); err != nil {
			log.Fatalf("Failed to add peer: %s due to :%v", p.Description, err)
		}
	}

	go watchBgpUpdates(bgpSrv)
}

// watchBgpUpdates listens for BGP updates and processes paths using a callback.
func watchBgpUpdates(bgpServer *bgpserver.Server) {
	// Define the callback for handling path events
	pathWatch := createPathWatchCallback()

	// Set up the watch event request
	watchRequest := &apipb.WatchEventRequest{
		Table: &apipb.WatchEventRequest_Table{
			Filters: []*apipb.WatchEventRequest_Table_Filter{
				{
					Type: apipb.WatchEventRequest_Table_Filter_BEST,
				},
			},
		},
	}

	// Watch for BGP events using the pathWatch callback
	if err := bgpServer.WatchEvent(context.Background(), watchRequest, pathWatch); err != nil {
		log.Printf("Failed to watch BGP events: %v", err)
	}
}

// createPathWatchCallback returns a callback function for handling path events.
func createPathWatchCallback() func(*apipb.WatchEventResponse) {
	return func(r *apipb.WatchEventResponse) {
		if table := r.GetTable(); table != nil {
			processPaths(table.Paths)
		}
	}
}

// processPaths iterates over paths and handles each one.
func processPaths(paths []*apipb.Path) {
	for _, path := range paths {
		if err := bgputil.HandlePath(path); err != nil {
			log.Printf("Failed to handle path: %v, error was %v", path, err)
		}
	}
}

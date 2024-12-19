// Package main is the entry point for the GoBGPNet application.
package main

import (
	"context"
	"io"
	"log"

	"github.com/karasz/gobgpnet/config"
	"github.com/karasz/gobgpnet/routes"

	apipb "github.com/osrg/gobgp/v3/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	_, client, conn := setupConnection("config.yaml")
	defer conn.Close()

	stream := setupRouteStream(client)

	processRouteUpdates(stream)
}

func setupConnection(configPath string) (*config.Config, apipb.GobgpApiClient, *grpc.ClientConn) {
	cfg := loadConfig(configPath)
	conn := createGRPCClient(cfg)
	client := apipb.NewGobgpApiClient(conn)
	return cfg, client, conn
}

func setupRouteStream(client apipb.GobgpApiClient) apipb.GobgpApi_WatchEventClient {
	log.Println("Starting route monitoring...")
	stream, err := client.WatchEvent(context.Background(), &apipb.WatchEventRequest{
		Table: &apipb.WatchEventRequest_Table{
			Filters: []*apipb.WatchEventRequest_Table_Filter{
				{
					Type: apipb.WatchEventRequest_Table_Filter_BEST,
					Init: true,
				},
			},
		},
	})
	if err != nil {
		log.Printf("Error creating route monitor stream: %v", err)
	}
	return stream
}

func processRouteUpdates(stream apipb.GobgpApi_WatchEventClient) {
	for {
		if done := handleRouteUpdate(stream); done {
			return
		}
	}
}

func handleRouteUpdate(stream apipb.GobgpApi_WatchEventClient) bool {
	resp, err := stream.Recv()
	if err != nil {
		return handleStreamError(err)
	}

	if resp.GetTable() != nil && len(resp.GetTable().Paths) > 0 {
		handleRoutePaths(resp.GetTable().Paths)
	}
	return false
}

func handleStreamError(err error) bool {
	if err == io.EOF {
		log.Println("Stream closed by server")
		return true
	}
	log.Printf("Error receiving from stream: %v", err)
	return false
}

func handleRoutePaths(paths []*apipb.Path) {
	if err := routes.UpdateLocalRoutes(paths); err != nil {
		log.Printf("Error updating routes: %v", err)
	}
}

func loadConfig(configPath string) *config.Config {
	cfg, err := config.Load(configPath)
	if err != nil {
		log.Printf("Error loading config: %v", err)
	}
	return cfg
}

func createGRPCClient(cfg *config.Config) *grpc.ClientConn {
	conn, err := grpc.NewClient(cfg.GoBGPServer, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Error creating gRPC client: %v", err)
	}
	return conn
}

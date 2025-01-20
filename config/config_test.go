package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Create a temporary file to simulate the configuration file
	tmpFile, err := os.CreateTemp("", "config_test_*.yaml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Write valid YAML content to the temporary file
	yamlContent := `
peers:
  - asn: 65001
	ip: "192.0.2.1"
	description: "Test Peer 1"
  - asn: 65002
	ip: "192.0.2.2"
	description: "Test Peer 2"
`
	if _, err := tmpFile.Write([]byte(yamlContent)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	_ = tmpFile.Close()

	// Call the Load function with the path to the temporary file
	cfg, err := Load(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to load config: %v", err)
	}

	// Verify that the returned Config object matches the expected values
	if len(cfg.Peers) != 2 {
		t.Errorf("Expected 2 peers, got %d", len(cfg.Peers))
	}
	if cfg.Peers[0].Asn != 65001 || cfg.Peers[0].IP != "192.0.2.1" || cfg.Peers[0].Description != "Test Peer 1" {
		t.Errorf("Unexpected peer 1: %+v", cfg.Peers[0])
	}
	if cfg.Peers[1].Asn != 65002 || cfg.Peers[1].IP != "192.0.2.2" || cfg.Peers[1].Description != "Test Peer 2" {
		t.Errorf("Unexpected peer 2: %+v", cfg.Peers[1])
	}
}

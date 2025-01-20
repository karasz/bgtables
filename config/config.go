// Package config provides functions to load and parse the configuration file.
package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"darvaza.org/slog"
	"gopkg.in/yaml.v2"

	"darvaza.org/slog/handlers/discard"
	"darvaza.org/x/config"
)

// Peer describes a BGP peer.
type Peer struct {
	Asn         uint32 `yaml:"asn"`
	IP          string `yaml:"ip"`
	Description string `yaml:"description"`
}

// Config describes the internal BGP Server.
type Config struct {
	Logger          slog.Logger
	Context         context.Context
	Peers           []Peer
	GracefulTimeout time.Duration
}

// Load loads the configuration from the given file path.
func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}
	_ = cfg.SetDefaults()

	return &cfg, nil
}

// SetDefaults fills gaps in the [Config].
func (sc *Config) SetDefaults() error {
	if sc.Context == nil {
		sc.Context = context.Background()
	}

	if sc.Logger == nil {
		sc.Logger = discard.New()
	}

	return config.Set(sc)
}

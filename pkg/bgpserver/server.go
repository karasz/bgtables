// Package bgpserver implements a server that has BGP client-server functionality
package bgpserver

import (
	"github.com/karasz/bgtables/config"
	"github.com/osrg/gobgp/v3/pkg/server"
)

// Server is a BGP server wrapping a [core.ErrGroup].
type Server struct {
	*server.BgpServer
	*config.Config
}

// New initializes and returns a new Server instance.
// It takes a configuration object and an error group as parameters.
// If the configuration is nil, a default configuration is created.
// The function sets default values for the configuration and
// initializes a new BGP server. Returns the Server instance or an error.
func New(sc *config.Config) (*Server, error) {
	if sc == nil {
		sc = new(config.Config)
	}

	if err := sc.SetDefaults(); err != nil {
		return nil, err
	}

	bgpsrv := server.NewBgpServer()

	srv := &Server{
		bgpsrv,
		sc,
	}
	return srv, nil
}

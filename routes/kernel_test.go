package routes

import (
	"testing"

	apipb "github.com/osrg/gobgp/v3/api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestUpdateLocalRoutes(t *testing.T) {
	prefix := &apipb.IPAddressPrefix{
		Prefix:    "192.168.1.0",
		PrefixLen: 24,
	}

	nlriAny, err := anypb.New(prefix)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name        string
		paths       []*apipb.Path
		expectError bool
	}{
		{
			name:        "Empty paths",
			paths:       []*apipb.Path{},
			expectError: false,
		},
		{
			name: "Valid paths",
			paths: []*apipb.Path{
				{
					Nlri: nlriAny,
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := UpdateLocalRoutes(tt.paths)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCreateRouteFromPath(t *testing.T) {
	prefix := &apipb.IPAddressPrefix{
		Prefix:    "192.168.1.0",
		PrefixLen: 24,
	}

	nlriAny, err := anypb.New(prefix)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name      string
		path      *apipb.Path
		expectNil bool
	}{
		{
			name: "Valid path",
			path: &apipb.Path{
				Nlri: nlriAny,
			},
			expectNil: false,
		},
		{
			name: "Invalid NLRI",
			path: &apipb.Path{
				Nlri: &anypb.Any{
					TypeUrl: "type.googleapis.com/apipb.IPAddressPrefix",
					Value:   []byte(`invalid`),
				},
			},
			expectNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := createRouteFromPath(tt.path)
			if tt.expectNil {
				assert.Nil(t, result)
			} else {
				assert.NotNil(t, result)
			}
		})
	}
}

/*
func TestUpdateRoute(t *testing.T) {
	tests := []struct {
		name        string
		cidr        string
		route       *netlink.Route
		expectError bool
	}{
		{
			name: "Valid route",
			cidr: "10.17.254.0/24",
			route: &netlink.Route{
				Dst: &net.IPNet{
					IP:   net.ParseIP("10.17.254.0"),
					Mask: net.CIDRMask(24, 32),
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := updateRoute(tt.cidr, tt.route)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
*/

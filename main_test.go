package main

import (
	"errors"
	"io"
	"testing"

	apipb "github.com/osrg/gobgp/v3/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"google.golang.org/protobuf/types/known/anypb"
)

// Mock implementations
type MockGobgpApiClient struct {
	mock.Mock
}

// Tests
func TestSetupConnection(t *testing.T) {
	tests := []struct {
		name        string
		configPath  string
		shouldError bool
	}{
		{
			name:        "Valid config path",
			configPath:  "testdata/config.yaml",
			shouldError: false,
		},
		{
			name:        "Invalid config path",
			configPath:  "testdata/nonexistent.yaml",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil && !tt.shouldError {
					t.Errorf("setupConnection() panicked unexpectedly: %v", r)
				}
			}()

			cfg, client, conn := setupConnection(tt.configPath)
			if !tt.shouldError {
				assert.NotNil(t, cfg)
				assert.NotNil(t, client)
				conn.Close()
			}
		})
	}
}

func TestHandleStreamError(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedDone bool
	}{
		{
			name:         "EOF error",
			err:          io.EOF,
			expectedDone: true,
		},
		{
			name:         "Other error",
			err:          errors.New("random error"),
			expectedDone: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			done := handleStreamError(tt.err)
			assert.Equal(t, tt.expectedDone, done)
		})
	}
}

func TestHandleRoutePaths(t *testing.T) {
	paths := []*apipb.Path{
		{
			Nlri: &anypb.Any{
				TypeUrl: "type.googleapis.com/apipb.IPAddressPrefix",
				Value:   []byte{192, 168, 1, 0},
			},
			Pattrs: []*anypb.Any{},
		},
	}

	handleRoutePaths(paths)
}

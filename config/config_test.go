package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//revive:disable:cognitive-complexity
func TestLoad(t *testing.T) {
	//revive:enable:cognitive-complexity
	// Test cases
	tests := []struct {
		name        string
		configYAML  string
		expectError bool
		expected    *Config
	}{
		{
			name: "Valid config",
			configYAML: `
gobgp_server: "localhost:50051"
`,
			expectError: false,
			expected: &Config{
				GoBGPServer: "localhost:50051",
			},
		},
		{
			name: "Invalid YAML",
			configYAML: `
gobgp_server: "localhost:50051
`,
			expectError: true,
			expected:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create temporary config file
			tmpfile, err := os.CreateTemp("", "config-*.yaml")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			// Write config content
			if _, err := tmpfile.Write([]byte(tt.configYAML)); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			// Test Load function
			config, err := Load(tmpfile.Name())
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, config)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected.GoBGPServer, config.GoBGPServer)
			}
		})
	}
}

func TestLoadNonExistentFile(t *testing.T) {
	_, err := Load("nonexistent.yaml")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to open config file")
}

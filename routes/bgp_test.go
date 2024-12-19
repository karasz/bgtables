package routes

import (
	"errors"
	"io"
	"testing"

	apipb "github.com/osrg/gobgp/v3/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestFetchRoutes(t *testing.T) {
	tests := []struct {
		name string

		setupMock   func(*MockGobgpApiClient, *MockListPathClient)
		expectError bool
		expectPaths int
	}{
		{
			name: "Successful fetch",
			setupMock: func(client *MockGobgpApiClient, stream *MockListPathClient) {
				stream.On("Recv").Return(&apipb.ListPathResponse{
					Destination: &apipb.Destination{
						Paths: []*apipb.Path{{}, {}},
					},
				}, nil).Once()
				stream.On("Recv").Return((*apipb.ListPathResponse)(nil), io.EOF)
				client.On("ListPath", mock.Anything, mock.Anything).Return(stream, nil)
			},
			expectError: false,
			expectPaths: 2,
		},
		{
			name: "ListPath error",
			setupMock: func(client *MockGobgpApiClient, stream *MockListPathClient) {
				client.On("ListPath", mock.Anything, mock.Anything).Return(nil, errors.New("connection error"))
			},
			expectError: true,
			expectPaths: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := new(MockGobgpApiClient)
			mockStream := new(MockListPathClient)
			tt.setupMock(mockClient, mockStream)

			paths, err := FetchRoutes(mockClient)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, paths, tt.expectPaths)
			}

			mockClient.AssertExpectations(t)
			mockStream.AssertExpectations(t)
		})
	}
}

func TestParseNlriToCIDR(t *testing.T) {
	tests := []struct {
		name        string
		input       *anypb.Any
		expected    string
		expectError bool
	}{
		{
			name: "Valid CIDR",
			input: func() *anypb.Any {
				prefix := &apipb.IPAddressPrefix{
					PrefixLen: 24,
					Prefix:    "192.168.1.0",
				}
				any, err := anypb.New(prefix)
				if err != nil {
					t.Fatalf("failed to create Any message: %v", err)
				}
				return any
			}(),
			expected:    "192.168.1.0/24",
			expectError: false,
		},
		{
			name: "Invalid CIDR",
			input: func() *anypb.Any {
				prefix := &apipb.IPAddressPrefix{
					PrefixLen: 0,
					Prefix:    "",
				}
				any, err := anypb.New(prefix)
				if err != nil {
					t.Fatalf("failed to create Any message: %v", err)
				}
				return any
			}(),
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseNlriToCIDR(tt.input)
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

package routes

import (
	"context"
	"errors"

	apipb "github.com/osrg/gobgp/v3/api"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Mock GobgpAPIClient
type MockGobgpAPIClient struct {
	mock.Mock
}

type MockListPathClient struct {
	mock.Mock
}

type MockAddBmpClient struct {
	mock.Mock
	grpc.ClientStream
}

func (m *MockListPathClient) Recv() (*apipb.ListPathResponse, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	resp, ok := args.Get(0).(*apipb.ListPathResponse)
	if !ok {
		return nil, errors.New("invalid response type")
	}
	return resp, args.Error(1)
}

func (*MockListPathClient) RecvMsg(any) error            { return nil }
func (*MockListPathClient) SendMsg(any) error            { return nil }
func (*MockListPathClient) Header() (metadata.MD, error) { return nil, nil }
func (*MockListPathClient) Trailer() metadata.MD         { return nil }
func (*MockListPathClient) CloseSend() error             { return nil }
func (*MockListPathClient) Context() context.Context     { return context.Background() }

func (m *MockGobgpAPIClient) ListPath(ctx context.Context, in *apipb.ListPathRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_ListPathClient, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	client, ok := args.Get(0).(apipb.GobgpApi_ListPathClient)
	if !ok {
		return nil, errors.New("invalid client type")
	}
	return client, args.Error(1)
}

func (*MockGobgpAPIClient) AddBmp(_ context.Context, _ *apipb.AddBmpRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) AddDefinedSet(_ context.Context, _ *apipb.AddDefinedSetRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) AddDynamicNeighbor(_ context.Context, _ *apipb.AddDynamicNeighborRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) AddPath(_ context.Context, _ *apipb.AddPathRequest,
	_ ...grpc.CallOption) (*apipb.AddPathResponse, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) AddPathStream(_ context.Context,
	_ ...grpc.CallOption) (apipb.GobgpApi_AddPathStreamClient, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) AddPeer(_ context.Context, _ *apipb.AddPeerRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) AddPeerGroup(_ context.Context, _ *apipb.AddPeerGroupRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) AddPolicy(_ context.Context, _ *apipb.AddPolicyRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) AddRpki(_ context.Context, _ *apipb.AddRpkiRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) AddStatement(_ context.Context, _ *apipb.AddStatementRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) AddVrf(_ context.Context, _ *apipb.AddVrfRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DeleteBmp(_ context.Context, _ *apipb.DeleteBmpRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DeleteDefinedSet(_ context.Context, _ *apipb.DeleteDefinedSetRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DeletePath(_ context.Context, _ *apipb.DeletePathRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DeletePeer(_ context.Context, _ *apipb.DeletePeerRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DeletePeerGroup(_ context.Context, _ *apipb.DeletePeerGroupRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DeletePolicy(_ context.Context, _ *apipb.DeletePolicyRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DeleteRpki(_ context.Context, _ *apipb.DeleteRpkiRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DeleteStatement(_ context.Context, _ *apipb.DeleteStatementRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DeleteVrf(_ context.Context, _ *apipb.DeleteVrfRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DisableMrt(_ context.Context, _ *apipb.DisableMrtRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DisablePeer(_ context.Context, _ *apipb.DisablePeerRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DisableRpki(_ context.Context, _ *apipb.DisableRpkiRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) EnableMrt(_ context.Context, _ *apipb.EnableMrtRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) EnablePeer(_ context.Context, _ *apipb.EnablePeerRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) EnableRpki(_ context.Context, _ *apipb.EnableRpkiRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) GetBgp(_ context.Context, _ *apipb.GetBgpRequest,
	_ ...grpc.CallOption) (*apipb.GetBgpResponse, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ResetPeer(_ context.Context, _ *apipb.ResetPeerRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ResetRpki(_ context.Context, _ *apipb.ResetRpkiRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) SetLogLevel(_ context.Context, _ *apipb.SetLogLevelRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ShutdownPeer(_ context.Context, _ *apipb.ShutdownPeerRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) AddPolicyAssignment(_ context.Context, _ *apipb.AddPolicyAssignmentRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DeleteDynamicNeighbor(_ context.Context, _ *apipb.DeleteDynamicNeighborRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) DeletePolicyAssignment(_ context.Context, _ *apipb.DeletePolicyAssignmentRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) EnableZebra(_ context.Context, _ *apipb.EnableZebraRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) GetTable(_ context.Context, _ *apipb.GetTableRequest,
	_ ...grpc.CallOption) (*apipb.GetTableResponse, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ListBmp(_ context.Context, _ *apipb.ListBmpRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_ListBmpClient, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ListDefinedSet(_ context.Context, _ *apipb.ListDefinedSetRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_ListDefinedSetClient, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ListDynamicNeighbor(_ context.Context, _ *apipb.ListDynamicNeighborRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_ListDynamicNeighborClient, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ListPeer(_ context.Context, _ *apipb.ListPeerRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_ListPeerClient, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ListPeerGroup(_ context.Context, _ *apipb.ListPeerGroupRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_ListPeerGroupClient, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ListPolicy(_ context.Context, _ *apipb.ListPolicyRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_ListPolicyClient, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ListPolicyAssignment(_ context.Context, _ *apipb.ListPolicyAssignmentRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_ListPolicyAssignmentClient, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ListRpki(_ context.Context, _ *apipb.ListRpkiRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_ListRpkiClient, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ListRpkiTable(_ context.Context, _ *apipb.ListRpkiTableRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_ListRpkiTableClient, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ListStatement(_ context.Context, _ *apipb.ListStatementRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_ListStatementClient, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) ListVrf(_ context.Context, _ *apipb.ListVrfRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_ListVrfClient, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) SetPolicies(_ context.Context, _ *apipb.SetPoliciesRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) SetPolicyAssignment(_ context.Context, _ *apipb.SetPolicyAssignmentRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) StartBgp(_ context.Context, _ *apipb.StartBgpRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) StopBgp(_ context.Context, _ *apipb.StopBgpRequest,
	_ ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) UpdatePeer(_ context.Context, _ *apipb.UpdatePeerRequest,
	_ ...grpc.CallOption) (*apipb.UpdatePeerResponse, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) UpdatePeerGroup(_ context.Context, _ *apipb.UpdatePeerGroupRequest,
	_ ...grpc.CallOption) (*apipb.UpdatePeerGroupResponse, error) {
	return nil, nil
}

func (*MockGobgpAPIClient) WatchEvent(_ context.Context, _ *apipb.WatchEventRequest,
	_ ...grpc.CallOption) (apipb.GobgpApi_WatchEventClient, error) {
	return nil, nil
}

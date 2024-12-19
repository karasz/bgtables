package routes

import (
	"context"

	apipb "github.com/osrg/gobgp/v3/api"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Mock GobgpApiClient
type MockGobgpApiClient struct {
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
	return args.Get(0).(*apipb.ListPathResponse), args.Error(1)
}

func (m *MockListPathClient) RecvMsg(interface{}) error    { return nil }
func (m *MockListPathClient) SendMsg(interface{}) error    { return nil }
func (m *MockListPathClient) Header() (metadata.MD, error) { return nil, nil }
func (m *MockListPathClient) Trailer() metadata.MD         { return nil }
func (m *MockListPathClient) CloseSend() error             { return nil }
func (m *MockListPathClient) Context() context.Context     { return context.Background() }

func (m *MockGobgpApiClient) ListPath(ctx context.Context, in *apipb.ListPathRequest, opts ...grpc.CallOption) (apipb.GobgpApi_ListPathClient, error) {
	args := m.Called(ctx, in)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(apipb.GobgpApi_ListPathClient), args.Error(1)
}

func (m *MockGobgpApiClient) AddBmp(ctx context.Context, in *apipb.AddBmpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) AddDefinedSet(ctx context.Context, in *apipb.AddDefinedSetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) AddDynamicNeighbor(ctx context.Context, in *apipb.AddDynamicNeighborRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) AddPath(ctx context.Context, in *apipb.AddPathRequest, opts ...grpc.CallOption) (*apipb.AddPathResponse, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) AddPathStream(ctx context.Context, opts ...grpc.CallOption) (apipb.GobgpApi_AddPathStreamClient, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) AddPeer(ctx context.Context, in *apipb.AddPeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) AddPeerGroup(ctx context.Context, in *apipb.AddPeerGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) AddPolicy(ctx context.Context, in *apipb.AddPolicyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) AddRpki(ctx context.Context, in *apipb.AddRpkiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) AddStatement(ctx context.Context, in *apipb.AddStatementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) AddVrf(ctx context.Context, in *apipb.AddVrfRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DeleteBmp(ctx context.Context, in *apipb.DeleteBmpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DeleteDefinedSet(ctx context.Context, in *apipb.DeleteDefinedSetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DeletePath(ctx context.Context, in *apipb.DeletePathRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DeletePeer(ctx context.Context, in *apipb.DeletePeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DeletePeerGroup(ctx context.Context, in *apipb.DeletePeerGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DeletePolicy(ctx context.Context, in *apipb.DeletePolicyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DeleteRpki(ctx context.Context, in *apipb.DeleteRpkiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DeleteStatement(ctx context.Context, in *apipb.DeleteStatementRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DeleteVrf(ctx context.Context, in *apipb.DeleteVrfRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DisableMrt(ctx context.Context, in *apipb.DisableMrtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DisablePeer(ctx context.Context, in *apipb.DisablePeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DisableRpki(ctx context.Context, in *apipb.DisableRpkiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) EnableMrt(ctx context.Context, in *apipb.EnableMrtRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) EnablePeer(ctx context.Context, in *apipb.EnablePeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) EnableRpki(ctx context.Context, in *apipb.EnableRpkiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) GetBgp(ctx context.Context, in *apipb.GetBgpRequest, opts ...grpc.CallOption) (*apipb.GetBgpResponse, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ResetPeer(ctx context.Context, in *apipb.ResetPeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ResetRpki(ctx context.Context, in *apipb.ResetRpkiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) SetLogLevel(ctx context.Context, in *apipb.SetLogLevelRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ShutdownPeer(ctx context.Context, in *apipb.ShutdownPeerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) AddPolicyAssignment(ctx context.Context, in *apipb.AddPolicyAssignmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DeleteDynamicNeighbor(ctx context.Context, in *apipb.DeleteDynamicNeighborRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) DeletePolicyAssignment(ctx context.Context, in *apipb.DeletePolicyAssignmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) EnableZebra(ctx context.Context, in *apipb.EnableZebraRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) GetTable(ctx context.Context, in *apipb.GetTableRequest, opts ...grpc.CallOption) (*apipb.GetTableResponse, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ListBmp(ctx context.Context, in *apipb.ListBmpRequest, opts ...grpc.CallOption) (apipb.GobgpApi_ListBmpClient, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ListDefinedSet(ctx context.Context, in *apipb.ListDefinedSetRequest, opts ...grpc.CallOption) (apipb.GobgpApi_ListDefinedSetClient, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ListDynamicNeighbor(ctx context.Context, in *apipb.ListDynamicNeighborRequest, opts ...grpc.CallOption) (apipb.GobgpApi_ListDynamicNeighborClient, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ListPeer(ctx context.Context, in *apipb.ListPeerRequest, opts ...grpc.CallOption) (apipb.GobgpApi_ListPeerClient, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ListPeerGroup(ctx context.Context, in *apipb.ListPeerGroupRequest, opts ...grpc.CallOption) (apipb.GobgpApi_ListPeerGroupClient, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ListPolicy(ctx context.Context, in *apipb.ListPolicyRequest, opts ...grpc.CallOption) (apipb.GobgpApi_ListPolicyClient, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ListPolicyAssignment(ctx context.Context, in *apipb.ListPolicyAssignmentRequest, opts ...grpc.CallOption) (apipb.GobgpApi_ListPolicyAssignmentClient, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ListRpki(ctx context.Context, in *apipb.ListRpkiRequest, opts ...grpc.CallOption) (apipb.GobgpApi_ListRpkiClient, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ListRpkiTable(ctx context.Context, in *apipb.ListRpkiTableRequest, opts ...grpc.CallOption) (apipb.GobgpApi_ListRpkiTableClient, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ListStatement(ctx context.Context, in *apipb.ListStatementRequest, opts ...grpc.CallOption) (apipb.GobgpApi_ListStatementClient, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) ListVrf(ctx context.Context, in *apipb.ListVrfRequest, opts ...grpc.CallOption) (apipb.GobgpApi_ListVrfClient, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) SetPolicies(ctx context.Context, in *apipb.SetPoliciesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) SetPolicyAssignment(ctx context.Context, in *apipb.SetPolicyAssignmentRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) StartBgp(ctx context.Context, in *apipb.StartBgpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) StopBgp(ctx context.Context, in *apipb.StopBgpRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) UpdatePeer(ctx context.Context, in *apipb.UpdatePeerRequest, opts ...grpc.CallOption) (*apipb.UpdatePeerResponse, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) UpdatePeerGroup(ctx context.Context, in *apipb.UpdatePeerGroupRequest, opts ...grpc.CallOption) (*apipb.UpdatePeerGroupResponse, error) {
	return nil, nil
}

func (m *MockGobgpApiClient) WatchEvent(ctx context.Context, in *apipb.WatchEventRequest, opts ...grpc.CallOption) (apipb.GobgpApi_WatchEventClient, error) {
	return nil, nil
}

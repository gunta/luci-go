// Code generated by MockGen. DO NOT EDIT.
// Source: rpc.pb.go

package buildbucketpb

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockisBatchRequest_Request_Request is a mock of isBatchRequest_Request_Request interface
type MockisBatchRequest_Request_Request struct {
	ctrl     *gomock.Controller
	recorder *MockisBatchRequest_Request_RequestMockRecorder
}

// MockisBatchRequest_Request_RequestMockRecorder is the mock recorder for MockisBatchRequest_Request_Request
type MockisBatchRequest_Request_RequestMockRecorder struct {
	mock *MockisBatchRequest_Request_Request
}

// NewMockisBatchRequest_Request_Request creates a new mock instance
func NewMockisBatchRequest_Request_Request(ctrl *gomock.Controller) *MockisBatchRequest_Request_Request {
	mock := &MockisBatchRequest_Request_Request{ctrl: ctrl}
	mock.recorder = &MockisBatchRequest_Request_RequestMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisBatchRequest_Request_Request) EXPECT() *MockisBatchRequest_Request_RequestMockRecorder {
	return m.recorder
}

// isBatchRequest_Request_Request mocks base method
func (m *MockisBatchRequest_Request_Request) isBatchRequest_Request_Request() {
	m.ctrl.Call(m, "isBatchRequest_Request_Request")
}

// isBatchRequest_Request_Request indicates an expected call of isBatchRequest_Request_Request
func (mr *MockisBatchRequest_Request_RequestMockRecorder) isBatchRequest_Request_Request() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isBatchRequest_Request_Request", reflect.TypeOf((*MockisBatchRequest_Request_Request)(nil).isBatchRequest_Request_Request))
}

// MockisBatchResponse_Response_Response is a mock of isBatchResponse_Response_Response interface
type MockisBatchResponse_Response_Response struct {
	ctrl     *gomock.Controller
	recorder *MockisBatchResponse_Response_ResponseMockRecorder
}

// MockisBatchResponse_Response_ResponseMockRecorder is the mock recorder for MockisBatchResponse_Response_Response
type MockisBatchResponse_Response_ResponseMockRecorder struct {
	mock *MockisBatchResponse_Response_Response
}

// NewMockisBatchResponse_Response_Response creates a new mock instance
func NewMockisBatchResponse_Response_Response(ctrl *gomock.Controller) *MockisBatchResponse_Response_Response {
	mock := &MockisBatchResponse_Response_Response{ctrl: ctrl}
	mock.recorder = &MockisBatchResponse_Response_ResponseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisBatchResponse_Response_Response) EXPECT() *MockisBatchResponse_Response_ResponseMockRecorder {
	return m.recorder
}

// isBatchResponse_Response_Response mocks base method
func (m *MockisBatchResponse_Response_Response) isBatchResponse_Response_Response() {
	m.ctrl.Call(m, "isBatchResponse_Response_Response")
}

// isBatchResponse_Response_Response indicates an expected call of isBatchResponse_Response_Response
func (mr *MockisBatchResponse_Response_ResponseMockRecorder) isBatchResponse_Response_Response() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isBatchResponse_Response_Response", reflect.TypeOf((*MockisBatchResponse_Response_Response)(nil).isBatchResponse_Response_Response))
}

// MockBuildsClient is a mock of BuildsClient interface
type MockBuildsClient struct {
	ctrl     *gomock.Controller
	recorder *MockBuildsClientMockRecorder
}

// MockBuildsClientMockRecorder is the mock recorder for MockBuildsClient
type MockBuildsClientMockRecorder struct {
	mock *MockBuildsClient
}

// NewMockBuildsClient creates a new mock instance
func NewMockBuildsClient(ctrl *gomock.Controller) *MockBuildsClient {
	mock := &MockBuildsClient{ctrl: ctrl}
	mock.recorder = &MockBuildsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuildsClient) EXPECT() *MockBuildsClientMockRecorder {
	return m.recorder
}

// GetBuild mocks base method
func (m *MockBuildsClient) GetBuild(ctx context.Context, in *GetBuildRequest, opts ...grpc.CallOption) (*Build, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBuild", varargs...)
	ret0, _ := ret[0].(*Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBuild indicates an expected call of GetBuild
func (mr *MockBuildsClientMockRecorder) GetBuild(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuild", reflect.TypeOf((*MockBuildsClient)(nil).GetBuild), varargs...)
}

// SearchBuilds mocks base method
func (m *MockBuildsClient) SearchBuilds(ctx context.Context, in *SearchBuildsRequest, opts ...grpc.CallOption) (*SearchBuildsResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SearchBuilds", varargs...)
	ret0, _ := ret[0].(*SearchBuildsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchBuilds indicates an expected call of SearchBuilds
func (mr *MockBuildsClientMockRecorder) SearchBuilds(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchBuilds", reflect.TypeOf((*MockBuildsClient)(nil).SearchBuilds), varargs...)
}

// Batch mocks base method
func (m *MockBuildsClient) Batch(ctx context.Context, in *BatchRequest, opts ...grpc.CallOption) (*BatchResponse, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Batch", varargs...)
	ret0, _ := ret[0].(*BatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Batch indicates an expected call of Batch
func (mr *MockBuildsClientMockRecorder) Batch(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Batch", reflect.TypeOf((*MockBuildsClient)(nil).Batch), varargs...)
}

// UpdateBuild mocks base method
func (m *MockBuildsClient) UpdateBuild(ctx context.Context, in *UpdateBuildRequest, opts ...grpc.CallOption) (*Build, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateBuild", varargs...)
	ret0, _ := ret[0].(*Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBuild indicates an expected call of UpdateBuild
func (mr *MockBuildsClientMockRecorder) UpdateBuild(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBuild", reflect.TypeOf((*MockBuildsClient)(nil).UpdateBuild), varargs...)
}

// MockBuildsServer is a mock of BuildsServer interface
type MockBuildsServer struct {
	ctrl     *gomock.Controller
	recorder *MockBuildsServerMockRecorder
}

// MockBuildsServerMockRecorder is the mock recorder for MockBuildsServer
type MockBuildsServerMockRecorder struct {
	mock *MockBuildsServer
}

// NewMockBuildsServer creates a new mock instance
func NewMockBuildsServer(ctrl *gomock.Controller) *MockBuildsServer {
	mock := &MockBuildsServer{ctrl: ctrl}
	mock.recorder = &MockBuildsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuildsServer) EXPECT() *MockBuildsServerMockRecorder {
	return m.recorder
}

// GetBuild mocks base method
func (m *MockBuildsServer) GetBuild(arg0 context.Context, arg1 *GetBuildRequest) (*Build, error) {
	ret := m.ctrl.Call(m, "GetBuild", arg0, arg1)
	ret0, _ := ret[0].(*Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBuild indicates an expected call of GetBuild
func (mr *MockBuildsServerMockRecorder) GetBuild(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuild", reflect.TypeOf((*MockBuildsServer)(nil).GetBuild), arg0, arg1)
}

// SearchBuilds mocks base method
func (m *MockBuildsServer) SearchBuilds(arg0 context.Context, arg1 *SearchBuildsRequest) (*SearchBuildsResponse, error) {
	ret := m.ctrl.Call(m, "SearchBuilds", arg0, arg1)
	ret0, _ := ret[0].(*SearchBuildsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchBuilds indicates an expected call of SearchBuilds
func (mr *MockBuildsServerMockRecorder) SearchBuilds(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchBuilds", reflect.TypeOf((*MockBuildsServer)(nil).SearchBuilds), arg0, arg1)
}

// Batch mocks base method
func (m *MockBuildsServer) Batch(arg0 context.Context, arg1 *BatchRequest) (*BatchResponse, error) {
	ret := m.ctrl.Call(m, "Batch", arg0, arg1)
	ret0, _ := ret[0].(*BatchResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Batch indicates an expected call of Batch
func (mr *MockBuildsServerMockRecorder) Batch(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Batch", reflect.TypeOf((*MockBuildsServer)(nil).Batch), arg0, arg1)
}

// UpdateBuild mocks base method
func (m *MockBuildsServer) UpdateBuild(arg0 context.Context, arg1 *UpdateBuildRequest) (*Build, error) {
	ret := m.ctrl.Call(m, "UpdateBuild", arg0, arg1)
	ret0, _ := ret[0].(*Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBuild indicates an expected call of UpdateBuild
func (mr *MockBuildsServerMockRecorder) UpdateBuild(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBuild", reflect.TypeOf((*MockBuildsServer)(nil).UpdateBuild), arg0, arg1)
}
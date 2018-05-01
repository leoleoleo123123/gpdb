// Code generated by MockGen. DO NOT EDIT.
// Source: idl/hub_to_agent.pb.go

// Package mock_idl is a generated GoMock package.
package mock_idl

import (
	"github.com/golang/mock/gomock"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gp_upgrade/idl"
	"reflect"
)

// MockAgentClient is a mock of AgentClient interface
type MockAgentClient struct {
	ctrl     *gomock.Controller
	recorder *MockAgentClientMockRecorder
}

// MockAgentClientMockRecorder is the mock recorder for MockAgentClient
type MockAgentClientMockRecorder struct {
	mock *MockAgentClient
}

// NewMockAgentClient creates a new mock instance
func NewMockAgentClient(ctrl *gomock.Controller) *MockAgentClient {
	mock := &MockAgentClient{ctrl: ctrl}
	mock.recorder = &MockAgentClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAgentClient) EXPECT() *MockAgentClientMockRecorder {
	return m.recorder
}

// CheckUpgradeStatus mocks base method
func (m *MockAgentClient) CheckUpgradeStatus(ctx context.Context, in *idl.CheckUpgradeStatusRequest, opts ...grpc.CallOption) (*idl.CheckUpgradeStatusReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckUpgradeStatus", varargs...)
	ret0, _ := ret[0].(*idl.CheckUpgradeStatusReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUpgradeStatus indicates an expected call of CheckUpgradeStatus
func (mr *MockAgentClientMockRecorder) CheckUpgradeStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUpgradeStatus", reflect.TypeOf((*MockAgentClient)(nil).CheckUpgradeStatus), varargs...)
}

// CheckConversionStatus mocks base method
func (m *MockAgentClient) CheckConversionStatus(ctx context.Context, in *idl.CheckConversionStatusRequest, opts ...grpc.CallOption) (*idl.CheckConversionStatusReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckConversionStatus", varargs...)
	ret0, _ := ret[0].(*idl.CheckConversionStatusReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckConversionStatus indicates an expected call of CheckConversionStatus
func (mr *MockAgentClientMockRecorder) CheckConversionStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckConversionStatus", reflect.TypeOf((*MockAgentClient)(nil).CheckConversionStatus), varargs...)
}

// CheckDiskUsageOnAgents mocks base method
func (m *MockAgentClient) CheckDiskUsageOnAgents(ctx context.Context, in *idl.CheckDiskUsageRequestToAgent, opts ...grpc.CallOption) (*idl.CheckDiskUsageReplyFromAgent, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckDiskUsageOnAgents", varargs...)
	ret0, _ := ret[0].(*idl.CheckDiskUsageReplyFromAgent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDiskUsageOnAgents indicates an expected call of CheckDiskUsageOnAgents
func (mr *MockAgentClientMockRecorder) CheckDiskUsageOnAgents(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDiskUsageOnAgents", reflect.TypeOf((*MockAgentClient)(nil).CheckDiskUsageOnAgents), varargs...)
}

// PingAgents mocks base method
func (m *MockAgentClient) PingAgents(ctx context.Context, in *idl.PingAgentsRequest, opts ...grpc.CallOption) (*idl.PingAgentsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PingAgents", varargs...)
	ret0, _ := ret[0].(*idl.PingAgentsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PingAgents indicates an expected call of PingAgents
func (mr *MockAgentClientMockRecorder) PingAgents(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingAgents", reflect.TypeOf((*MockAgentClient)(nil).PingAgents), varargs...)
}

// UpgradeConvertPrimarySegments mocks base method
func (m *MockAgentClient) UpgradeConvertPrimarySegments(ctx context.Context, in *idl.UpgradeConvertPrimarySegmentsRequest, opts ...grpc.CallOption) (*idl.UpgradeConvertPrimarySegmentsReply, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpgradeConvertPrimarySegments", varargs...)
	ret0, _ := ret[0].(*idl.UpgradeConvertPrimarySegmentsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeConvertPrimarySegments indicates an expected call of UpgradeConvertPrimarySegments
func (mr *MockAgentClientMockRecorder) UpgradeConvertPrimarySegments(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeConvertPrimarySegments", reflect.TypeOf((*MockAgentClient)(nil).UpgradeConvertPrimarySegments), varargs...)
}

// MockAgentServer is a mock of AgentServer interface
type MockAgentServer struct {
	ctrl     *gomock.Controller
	recorder *MockAgentServerMockRecorder
}

// MockAgentServerMockRecorder is the mock recorder for MockAgentServer
type MockAgentServerMockRecorder struct {
	mock *MockAgentServer
}

// NewMockAgentServer creates a new mock instance
func NewMockAgentServer(ctrl *gomock.Controller) *MockAgentServer {
	mock := &MockAgentServer{ctrl: ctrl}
	mock.recorder = &MockAgentServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAgentServer) EXPECT() *MockAgentServerMockRecorder {
	return m.recorder
}

// CheckUpgradeStatus mocks base method
func (m *MockAgentServer) CheckUpgradeStatus(arg0 context.Context, arg1 *idl.CheckUpgradeStatusRequest) (*idl.CheckUpgradeStatusReply, error) {
	ret := m.ctrl.Call(m, "CheckUpgradeStatus", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckUpgradeStatusReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUpgradeStatus indicates an expected call of CheckUpgradeStatus
func (mr *MockAgentServerMockRecorder) CheckUpgradeStatus(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUpgradeStatus", reflect.TypeOf((*MockAgentServer)(nil).CheckUpgradeStatus), arg0, arg1)
}

// CheckConversionStatus mocks base method
func (m *MockAgentServer) CheckConversionStatus(arg0 context.Context, arg1 *idl.CheckConversionStatusRequest) (*idl.CheckConversionStatusReply, error) {
	ret := m.ctrl.Call(m, "CheckConversionStatus", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckConversionStatusReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckConversionStatus indicates an expected call of CheckConversionStatus
func (mr *MockAgentServerMockRecorder) CheckConversionStatus(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckConversionStatus", reflect.TypeOf((*MockAgentServer)(nil).CheckConversionStatus), arg0, arg1)
}

// CheckDiskUsageOnAgents mocks base method
func (m *MockAgentServer) CheckDiskUsageOnAgents(arg0 context.Context, arg1 *idl.CheckDiskUsageRequestToAgent) (*idl.CheckDiskUsageReplyFromAgent, error) {
	ret := m.ctrl.Call(m, "CheckDiskUsageOnAgents", arg0, arg1)
	ret0, _ := ret[0].(*idl.CheckDiskUsageReplyFromAgent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckDiskUsageOnAgents indicates an expected call of CheckDiskUsageOnAgents
func (mr *MockAgentServerMockRecorder) CheckDiskUsageOnAgents(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckDiskUsageOnAgents", reflect.TypeOf((*MockAgentServer)(nil).CheckDiskUsageOnAgents), arg0, arg1)
}

// PingAgents mocks base method
func (m *MockAgentServer) PingAgents(arg0 context.Context, arg1 *idl.PingAgentsRequest) (*idl.PingAgentsReply, error) {
	ret := m.ctrl.Call(m, "PingAgents", arg0, arg1)
	ret0, _ := ret[0].(*idl.PingAgentsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PingAgents indicates an expected call of PingAgents
func (mr *MockAgentServerMockRecorder) PingAgents(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingAgents", reflect.TypeOf((*MockAgentServer)(nil).PingAgents), arg0, arg1)
}

// UpgradeConvertPrimarySegments mocks base method
func (m *MockAgentServer) UpgradeConvertPrimarySegments(arg0 context.Context, arg1 *idl.UpgradeConvertPrimarySegmentsRequest) (*idl.UpgradeConvertPrimarySegmentsReply, error) {
	ret := m.ctrl.Call(m, "UpgradeConvertPrimarySegments", arg0, arg1)
	ret0, _ := ret[0].(*idl.UpgradeConvertPrimarySegmentsReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeConvertPrimarySegments indicates an expected call of UpgradeConvertPrimarySegments
func (mr *MockAgentServerMockRecorder) UpgradeConvertPrimarySegments(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeConvertPrimarySegments", reflect.TypeOf((*MockAgentServer)(nil).UpgradeConvertPrimarySegments), arg0, arg1)
}
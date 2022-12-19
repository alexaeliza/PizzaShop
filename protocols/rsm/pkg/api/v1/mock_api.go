// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/atomix/multi-raft-storage/api/atomix/multiraft/v1 (interfaces: SessionServer,PartitionServer)

// Package client is a generated GoMock package.
package v1

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSessionServer is a mock of SessionServer interface.
type MockSessionServer struct {
	ctrl     *gomock.Controller
	recorder *MockSessionServerMockRecorder
}

// MockSessionServerMockRecorder is the mock recorder for MockSessionServer.
type MockSessionServerMockRecorder struct {
	mock *MockSessionServer
}

// NewMockSessionServer creates a new mock instance.
func NewMockSessionServer(ctrl *gomock.Controller) *MockSessionServer {
	mock := &MockSessionServer{ctrl: ctrl}
	mock.recorder = &MockSessionServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionServer) EXPECT() *MockSessionServerMockRecorder {
	return m.recorder
}

// ClosePrimitive mocks base method.
func (m *MockSessionServer) ClosePrimitive(arg0 context.Context, arg1 *ClosePrimitiveRequest) (*ClosePrimitiveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClosePrimitive", arg0, arg1)
	ret0, _ := ret[0].(*ClosePrimitiveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClosePrimitive indicates an expected call of ClosePrimitive.
func (mr *MockSessionServerMockRecorder) ClosePrimitive(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClosePrimitive", reflect.TypeOf((*MockSessionServer)(nil).ClosePrimitive), arg0, arg1)
}

// CreatePrimitive mocks base method.
func (m *MockSessionServer) CreatePrimitive(arg0 context.Context, arg1 *CreatePrimitiveRequest) (*CreatePrimitiveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePrimitive", arg0, arg1)
	ret0, _ := ret[0].(*CreatePrimitiveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePrimitive indicates an expected call of CreatePrimitive.
func (mr *MockSessionServerMockRecorder) CreatePrimitive(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePrimitive", reflect.TypeOf((*MockSessionServer)(nil).CreatePrimitive), arg0, arg1)
}

// MockPartitionServer is a mock of PartitionServer interface.
type MockPartitionServer struct {
	ctrl     *gomock.Controller
	recorder *MockPartitionServerMockRecorder
}

// MockPartitionServerMockRecorder is the mock recorder for MockPartitionServer.
type MockPartitionServerMockRecorder struct {
	mock *MockPartitionServer
}

// NewMockPartitionServer creates a new mock instance.
func NewMockPartitionServer(ctrl *gomock.Controller) *MockPartitionServer {
	mock := &MockPartitionServer{ctrl: ctrl}
	mock.recorder = &MockPartitionServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPartitionServer) EXPECT() *MockPartitionServerMockRecorder {
	return m.recorder
}

// CloseSession mocks base method.
func (m *MockPartitionServer) CloseSession(arg0 context.Context, arg1 *CloseSessionRequest) (*CloseSessionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSession", arg0, arg1)
	ret0, _ := ret[0].(*CloseSessionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseSession indicates an expected call of CloseSession.
func (mr *MockPartitionServerMockRecorder) CloseSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSession", reflect.TypeOf((*MockPartitionServer)(nil).CloseSession), arg0, arg1)
}

// KeepAlive mocks base method.
func (m *MockPartitionServer) KeepAlive(arg0 context.Context, arg1 *KeepAliveRequest) (*KeepAliveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeepAlive", arg0, arg1)
	ret0, _ := ret[0].(*KeepAliveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KeepAlive indicates an expected call of KeepAlive.
func (mr *MockPartitionServerMockRecorder) KeepAlive(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeepAlive", reflect.TypeOf((*MockPartitionServer)(nil).KeepAlive), arg0, arg1)
}

// OpenSession mocks base method.
func (m *MockPartitionServer) OpenSession(arg0 context.Context, arg1 *OpenSessionRequest) (*OpenSessionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenSession", arg0, arg1)
	ret0, _ := ret[0].(*OpenSessionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenSession indicates an expected call of OpenSession.
func (mr *MockPartitionServerMockRecorder) OpenSession(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenSession", reflect.TypeOf((*MockPartitionServer)(nil).OpenSession), arg0, arg1)
}

// MockTestServer is a mock of TestServer interface.
type MockTestServer struct {
	ctrl     *gomock.Controller
	recorder *MockTestServerMockRecorder
}

// MockTestServerMockRecorder is the mock recorder for MockTestServer.
type MockTestServerMockRecorder struct {
	mock *MockTestServer
}

// NewMockTestServer creates a new mock instance.
func NewMockTestServer(ctrl *gomock.Controller) *MockTestServer {
	mock := &MockTestServer{ctrl: ctrl}
	mock.recorder = &MockTestServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTestServer) EXPECT() *MockTestServerMockRecorder {
	return m.recorder
}

// TestProposal mocks base method.
func (m *MockTestServer) TestPropose(arg0 context.Context, arg1 *TestProposalRequest) (*TestProposalResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestPropose", arg0, arg1)
	ret0, _ := ret[0].(*TestProposalResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TestPropose indicates an expected call of TestPropose.
func (mr *MockTestServerMockRecorder) TestPropose(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestPropose", reflect.TypeOf((*MockTestServer)(nil).TestPropose), arg0, arg1)
}

// TestQuery mocks base method.
func (m *MockTestServer) TestQuery(arg0 context.Context, arg1 *TestQueryRequest) (*TestQueryResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestQuery", arg0, arg1)
	ret0, _ := ret[0].(*TestQueryResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TestQuery indicates an expected call of TestQuery.
func (mr *MockTestServerMockRecorder) TestQuery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestQuery", reflect.TypeOf((*MockTestServer)(nil).TestQuery), arg0, arg1)
}

// TestStreamPropose mocks base method.
func (m *MockTestServer) TestStreamPropose(arg0 *TestProposalRequest, arg1 Test_TestStreamProposeServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestStreamPropose", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TestStreamPropose indicates an expected call of TestStreamPropose.
func (mr *MockTestServerMockRecorder) TestStreamPropose(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestStreamPropose", reflect.TypeOf((*MockTestServer)(nil).TestStreamPropose), arg0, arg1)
}

// TestStreamQuery mocks base method.
func (m *MockTestServer) TestStreamQuery(arg0 *TestQueryRequest, arg1 Test_TestStreamQueryServer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestStreamQuery", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TestStreamQuery indicates an expected call of TestStreamQuery.
func (mr *MockTestServerMockRecorder) TestStreamQuery(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestStreamQuery", reflect.TypeOf((*MockTestServer)(nil).TestStreamQuery), arg0, arg1)
}

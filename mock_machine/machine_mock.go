
// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/toransahu/grpc-eg-go/machine (interfaces: MachineClient,Machine_ExecuteServer,Machine_ExecuteClient)

// Package mock_machine is a generated GoMock package.
package mock_machine

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	machine "github.com/toransahu/grpc-eg-go/machine"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockMachineClient is a mock of MachineClient interface
type MockMachineClient struct {
	ctrl     *gomock.Controller
	recorder *MockMachineClientMockRecorder
}

// MockMachineClientMockRecorder is the mock recorder for MockMachineClient
type MockMachineClientMockRecorder struct {
	mock *MockMachineClient
}

// NewMockMachineClient creates a new mock instance
func NewMockMachineClient(ctrl *gomock.Controller) *MockMachineClient {
	mock := &MockMachineClient{ctrl: ctrl}
	mock.recorder = &MockMachineClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMachineClient) EXPECT() *MockMachineClientMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockMachineClient) Execute(arg0 context.Context, arg1 ...grpc.CallOption) (machine.Machine_ExecuteClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Execute", varargs...)
	ret0, _ := ret[0].(machine.Machine_ExecuteClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockMachineClientMockRecorder) Execute(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockMachineClient)(nil).Execute), varargs...)
}

// MockMachine_ExecuteServer is a mock of Machine_ExecuteServer interface
type MockMachine_ExecuteServer struct {
	ctrl     *gomock.Controller
	recorder *MockMachine_ExecuteServerMockRecorder
}

// MockMachine_ExecuteServerMockRecorder is the mock recorder for MockMachine_ExecuteServer
type MockMachine_ExecuteServerMockRecorder struct {
	mock *MockMachine_ExecuteServer
}

// NewMockMachine_ExecuteServer creates a new mock instance
func NewMockMachine_ExecuteServer(ctrl *gomock.Controller) *MockMachine_ExecuteServer {
	mock := &MockMachine_ExecuteServer{ctrl: ctrl}
	mock.recorder = &MockMachine_ExecuteServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMachine_ExecuteServer) EXPECT() *MockMachine_ExecuteServerMockRecorder {
	return m.recorder
}

// Context mocks base method
func (m *MockMachine_ExecuteServer) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockMachine_ExecuteServerMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockMachine_ExecuteServer)(nil).Context))
}

// Recv mocks base method
func (m *MockMachine_ExecuteServer) Recv() (*machine.Instruction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*machine.Instruction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockMachine_ExecuteServerMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockMachine_ExecuteServer)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockMachine_ExecuteServer) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockMachine_ExecuteServerMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockMachine_ExecuteServer)(nil).RecvMsg), arg0)
}

// Send mocks base method
func (m *MockMachine_ExecuteServer) Send(arg0 *machine.Result) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockMachine_ExecuteServerMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockMachine_ExecuteServer)(nil).Send), arg0)
}

// SendHeader mocks base method
func (m *MockMachine_ExecuteServer) SendHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendHeader indicates an expected call of SendHeader
func (mr *MockMachine_ExecuteServerMockRecorder) SendHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendHeader", reflect.TypeOf((*MockMachine_ExecuteServer)(nil).SendHeader), arg0)
}

// SendMsg mocks base method
func (m *MockMachine_ExecuteServer) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockMachine_ExecuteServerMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockMachine_ExecuteServer)(nil).SendMsg), arg0)
}

// SetHeader mocks base method
func (m *MockMachine_ExecuteServer) SetHeader(arg0 metadata.MD) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHeader", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHeader indicates an expected call of SetHeader
func (mr *MockMachine_ExecuteServerMockRecorder) SetHeader(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHeader", reflect.TypeOf((*MockMachine_ExecuteServer)(nil).SetHeader), arg0)
}

// SetTrailer mocks base method
func (m *MockMachine_ExecuteServer) SetTrailer(arg0 metadata.MD) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTrailer", arg0)
}

// SetTrailer indicates an expected call of SetTrailer
func (mr *MockMachine_ExecuteServerMockRecorder) SetTrailer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTrailer", reflect.TypeOf((*MockMachine_ExecuteServer)(nil).SetTrailer), arg0)
}

// MockMachine_ExecuteClient is a mock of Machine_ExecuteClient interface
type MockMachine_ExecuteClient struct {
	ctrl     *gomock.Controller
	recorder *MockMachine_ExecuteClientMockRecorder
}

// MockMachine_ExecuteClientMockRecorder is the mock recorder for MockMachine_ExecuteClient
type MockMachine_ExecuteClientMockRecorder struct {
	mock *MockMachine_ExecuteClient
}

// NewMockMachine_ExecuteClient creates a new mock instance
func NewMockMachine_ExecuteClient(ctrl *gomock.Controller) *MockMachine_ExecuteClient {
	mock := &MockMachine_ExecuteClient{ctrl: ctrl}
	mock.recorder = &MockMachine_ExecuteClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
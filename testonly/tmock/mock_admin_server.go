// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/google/trillian (interfaces: TrillianAdminServer)

// Package tmock is a generated GoMock package.
package tmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	trillian "github.com/google/trillian"
)

// MockTrillianAdminServer is a mock of TrillianAdminServer interface.
type MockTrillianAdminServer struct {
	ctrl     *gomock.Controller
	recorder *MockTrillianAdminServerMockRecorder
}

// MockTrillianAdminServerMockRecorder is the mock recorder for MockTrillianAdminServer.
type MockTrillianAdminServerMockRecorder struct {
	mock *MockTrillianAdminServer
}

// NewMockTrillianAdminServer creates a new mock instance.
func NewMockTrillianAdminServer(ctrl *gomock.Controller) *MockTrillianAdminServer {
	mock := &MockTrillianAdminServer{ctrl: ctrl}
	mock.recorder = &MockTrillianAdminServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrillianAdminServer) EXPECT() *MockTrillianAdminServerMockRecorder {
	return m.recorder
}

// CreateTree mocks base method.
func (m *MockTrillianAdminServer) CreateTree(arg0 context.Context, arg1 *trillian.CreateTreeRequest) (*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTree", arg0, arg1)
	ret0, _ := ret[0].(*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTree indicates an expected call of CreateTree.
func (mr *MockTrillianAdminServerMockRecorder) CreateTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTree", reflect.TypeOf((*MockTrillianAdminServer)(nil).CreateTree), arg0, arg1)
}

// DeleteTree mocks base method.
func (m *MockTrillianAdminServer) DeleteTree(arg0 context.Context, arg1 *trillian.DeleteTreeRequest) (*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTree", arg0, arg1)
	ret0, _ := ret[0].(*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTree indicates an expected call of DeleteTree.
func (mr *MockTrillianAdminServerMockRecorder) DeleteTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTree", reflect.TypeOf((*MockTrillianAdminServer)(nil).DeleteTree), arg0, arg1)
}

// GetTree mocks base method.
func (m *MockTrillianAdminServer) GetTree(arg0 context.Context, arg1 *trillian.GetTreeRequest) (*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTree", arg0, arg1)
	ret0, _ := ret[0].(*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTree indicates an expected call of GetTree.
func (mr *MockTrillianAdminServerMockRecorder) GetTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTree", reflect.TypeOf((*MockTrillianAdminServer)(nil).GetTree), arg0, arg1)
}

// ListTrees mocks base method.
func (m *MockTrillianAdminServer) ListTrees(arg0 context.Context, arg1 *trillian.ListTreesRequest) (*trillian.ListTreesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTrees", arg0, arg1)
	ret0, _ := ret[0].(*trillian.ListTreesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTrees indicates an expected call of ListTrees.
func (mr *MockTrillianAdminServerMockRecorder) ListTrees(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTrees", reflect.TypeOf((*MockTrillianAdminServer)(nil).ListTrees), arg0, arg1)
}

// UndeleteTree mocks base method.
func (m *MockTrillianAdminServer) UndeleteTree(arg0 context.Context, arg1 *trillian.UndeleteTreeRequest) (*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UndeleteTree", arg0, arg1)
	ret0, _ := ret[0].(*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UndeleteTree indicates an expected call of UndeleteTree.
func (mr *MockTrillianAdminServerMockRecorder) UndeleteTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UndeleteTree", reflect.TypeOf((*MockTrillianAdminServer)(nil).UndeleteTree), arg0, arg1)
}

// UpdateTree mocks base method.
func (m *MockTrillianAdminServer) UpdateTree(arg0 context.Context, arg1 *trillian.UpdateTreeRequest) (*trillian.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTree", arg0, arg1)
	ret0, _ := ret[0].(*trillian.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTree indicates an expected call of UpdateTree.
func (mr *MockTrillianAdminServerMockRecorder) UpdateTree(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTree", reflect.TypeOf((*MockTrillianAdminServer)(nil).UpdateTree), arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: ./repositoryv1/iface.go

// Package repositoryv1mock is a generated GoMock package.
package repositoryv1mock

import (
	proto "main/gen/go/proto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStaffV1Repository is a mock of StaffV1Repository interface.
type MockStaffV1Repository struct {
	ctrl     *gomock.Controller
	recorder *MockStaffV1RepositoryMockRecorder
}

// MockStaffV1RepositoryMockRecorder is the mock recorder for MockStaffV1Repository.
type MockStaffV1RepositoryMockRecorder struct {
	mock *MockStaffV1Repository
}

// NewMockStaffV1Repository creates a new mock instance.
func NewMockStaffV1Repository(ctrl *gomock.Controller) *MockStaffV1Repository {
	mock := &MockStaffV1Repository{ctrl: ctrl}
	mock.recorder = &MockStaffV1RepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStaffV1Repository) EXPECT() *MockStaffV1RepositoryMockRecorder {
	return m.recorder
}

// CountTotalStaff mocks base method.
func (m *MockStaffV1Repository) CountTotalStaff() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountTotalStaff")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountTotalStaff indicates an expected call of CountTotalStaff.
func (mr *MockStaffV1RepositoryMockRecorder) CountTotalStaff() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountTotalStaff", reflect.TypeOf((*MockStaffV1Repository)(nil).CountTotalStaff))
}

// CreateStaff mocks base method.
func (m *MockStaffV1Repository) CreateStaff(arg0 *proto.StaffV1) (*proto.StaffV1, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStaff", arg0)
	ret0, _ := ret[0].(*proto.StaffV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStaff indicates an expected call of CreateStaff.
func (mr *MockStaffV1RepositoryMockRecorder) CreateStaff(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStaff", reflect.TypeOf((*MockStaffV1Repository)(nil).CreateStaff), arg0)
}

// QueryAllStaffWithOffsetAndLimit mocks base method.
func (m *MockStaffV1Repository) QueryAllStaffWithOffsetAndLimit(offset, limit int) ([]*proto.StaffV1, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryAllStaffWithOffsetAndLimit", offset, limit)
	ret0, _ := ret[0].([]*proto.StaffV1)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryAllStaffWithOffsetAndLimit indicates an expected call of QueryAllStaffWithOffsetAndLimit.
func (mr *MockStaffV1RepositoryMockRecorder) QueryAllStaffWithOffsetAndLimit(offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryAllStaffWithOffsetAndLimit", reflect.TypeOf((*MockStaffV1Repository)(nil).QueryAllStaffWithOffsetAndLimit), offset, limit)
}
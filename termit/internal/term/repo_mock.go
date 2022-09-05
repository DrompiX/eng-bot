// Code generated by MockGen. DO NOT EDIT.
// Source: gitlab.ozon.dev/DrompiX/homework-2/termit/internal/term (interfaces: Repository)

// Package term is a generated GoMock package.
package term

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	user "gitlab.ozon.dev/DrompiX/homework-2/termit/internal/user"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AddTerm mocks base method.
func (m *MockRepository) AddTerm(arg0 context.Context, arg1 *Term) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTerm", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTerm indicates an expected call of AddTerm.
func (mr *MockRepositoryMockRecorder) AddTerm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTerm", reflect.TypeOf((*MockRepository)(nil).AddTerm), arg0, arg1)
}

// Atomic mocks base method.
func (m *MockRepository) Atomic(arg0 context.Context, arg1 func(Repository) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Atomic", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Atomic indicates an expected call of Atomic.
func (mr *MockRepositoryMockRecorder) Atomic(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Atomic", reflect.TypeOf((*MockRepository)(nil).Atomic), arg0, arg1)
}

// DeleteTerm mocks base method.
func (m *MockRepository) DeleteTerm(arg0 context.Context, arg1 *Term) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTerm", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTerm indicates an expected call of DeleteTerm.
func (mr *MockRepositoryMockRecorder) DeleteTerm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTerm", reflect.TypeOf((*MockRepository)(nil).DeleteTerm), arg0, arg1)
}

// GetAllTerms mocks base method.
func (m *MockRepository) GetAllTerms(arg0 context.Context, arg1 user.UserID) ([]*Term, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTerms", arg0, arg1)
	ret0, _ := ret[0].([]*Term)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTerms indicates an expected call of GetAllTerms.
func (mr *MockRepositoryMockRecorder) GetAllTerms(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTerms", reflect.TypeOf((*MockRepository)(nil).GetAllTerms), arg0, arg1)
}

// GetTermByName mocks base method.
func (m *MockRepository) GetTermByName(arg0 context.Context, arg1 user.UserID, arg2 string) (*Term, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTermByName", arg0, arg1, arg2)
	ret0, _ := ret[0].(*Term)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTermByName indicates an expected call of GetTermByName.
func (mr *MockRepositoryMockRecorder) GetTermByName(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTermByName", reflect.TypeOf((*MockRepository)(nil).GetTermByName), arg0, arg1, arg2)
}

// GetTermCount mocks base method.
func (m *MockRepository) GetTermCount(arg0 context.Context, arg1 user.UserID) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTermCount", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTermCount indicates an expected call of GetTermCount.
func (mr *MockRepositoryMockRecorder) GetTermCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTermCount", reflect.TypeOf((*MockRepository)(nil).GetTermCount), arg0, arg1)
}

// UpdateTerm mocks base method.
func (m *MockRepository) UpdateTerm(arg0 context.Context, arg1 *Term) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTerm", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTerm indicates an expected call of UpdateTerm.
func (mr *MockRepositoryMockRecorder) UpdateTerm(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTerm", reflect.TypeOf((*MockRepository)(nil).UpdateTerm), arg0, arg1)
}

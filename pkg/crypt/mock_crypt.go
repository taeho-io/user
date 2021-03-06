// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/taeho-io/user/pkg/crypt (interfaces: Crypt)

// Package crypt is a generated GoMock package.
package crypt

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCrypt is a mock of Crypt interface
type MockCrypt struct {
	ctrl     *gomock.Controller
	recorder *MockCryptMockRecorder
}

// MockCryptMockRecorder is the mock recorder for MockCrypt
type MockCryptMockRecorder struct {
	mock *MockCrypt
}

// NewMockCrypt creates a new mock instance
func NewMockCrypt(ctrl *gomock.Controller) *MockCrypt {
	mock := &MockCrypt{ctrl: ctrl}
	mock.recorder = &MockCryptMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCrypt) EXPECT() *MockCryptMockRecorder {
	return m.recorder
}

// HashPassword mocks base method
func (m *MockCrypt) HashPassword(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword
func (mr *MockCryptMockRecorder) HashPassword(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockCrypt)(nil).HashPassword), arg0)
}

// IsValidPassword mocks base method
func (m *MockCrypt) IsValidPassword(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidPassword", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidPassword indicates an expected call of IsValidPassword
func (mr *MockCryptMockRecorder) IsValidPassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidPassword", reflect.TypeOf((*MockCrypt)(nil).IsValidPassword), arg0, arg1)
}

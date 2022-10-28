// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/interfaces/interface.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockpersona is a mock of persona interface.
type Mockpersona struct {
	ctrl     *gomock.Controller
	recorder *MockpersonaMockRecorder
}

// MockpersonaMockRecorder is the mock recorder for Mockpersona.
type MockpersonaMockRecorder struct {
	mock *Mockpersona
}

// NewMockpersona creates a new mock instance.
func NewMockpersona(ctrl *gomock.Controller) *Mockpersona {
	mock := &Mockpersona{ctrl: ctrl}
	mock.recorder = &MockpersonaMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockpersona) EXPECT() *MockpersonaMockRecorder {
	return m.recorder
}

// FullName mocks base method.
func (m *Mockpersona) FullName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FullName")
	ret0, _ := ret[0].(string)
	return ret0
}

// FullName indicates an expected call of FullName.
func (mr *MockpersonaMockRecorder) FullName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FullName", reflect.TypeOf((*Mockpersona)(nil).FullName))
}

// GetAge mocks base method.
func (m *Mockpersona) GetAge() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAge")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetAge indicates an expected call of GetAge.
func (mr *MockpersonaMockRecorder) GetAge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAge", reflect.TypeOf((*Mockpersona)(nil).GetAge))
}

// IsMinor mocks base method.
func (m *Mockpersona) IsMinor() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMinor")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMinor indicates an expected call of IsMinor.
func (mr *MockpersonaMockRecorder) IsMinor() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMinor", reflect.TypeOf((*Mockpersona)(nil).IsMinor))
}
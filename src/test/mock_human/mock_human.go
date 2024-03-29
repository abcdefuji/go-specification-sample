// Code generated by MockGen. DO NOT EDIT.
// Source: human.go

// Package mock_test is a generated GoMock package.
package mock_human

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHuman is a mock of Human interface
type MockHuman struct {
	ctrl     *gomock.Controller
	recorder *MockHumanMockRecorder
}

// MockHumanMockRecorder is the mock recorder for MockHuman
type MockHumanMockRecorder struct {
	mock *MockHuman
}

// NewMockHuman creates a new mock instance
func NewMockHuman(ctrl *gomock.Controller) *MockHuman {
	mock := &MockHuman{ctrl: ctrl}
	mock.recorder = &MockHumanMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHuman) EXPECT() *MockHumanMockRecorder {
	return m.recorder
}

// Greeting mocks base method
func (m *MockHuman) Greeting(x string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Greeting", x)
	ret0, _ := ret[0].(string)
	return ret0
}

// Greeting indicates an expected call of Greeting
func (mr *MockHumanMockRecorder) Greeting(x interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Greeting", reflect.TypeOf((*MockHuman)(nil).Greeting), x)
}

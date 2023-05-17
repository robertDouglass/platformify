// Code generated by MockGen. DO NOT EDIT.
// Source: platformifier.go

// Package platformifier is a generated GoMock package.
package platformifier

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockplatformifier is a mock of platformifier interface.
type Mockplatformifier struct {
	ctrl     *gomock.Controller
	recorder *MockplatformifierMockRecorder
}

// MockplatformifierMockRecorder is the mock recorder for Mockplatformifier.
type MockplatformifierMockRecorder struct {
	mock *Mockplatformifier
}

// NewMockplatformifier creates a new mock instance.
func NewMockplatformifier(ctrl *gomock.Controller) *Mockplatformifier {
	mock := &Mockplatformifier{ctrl: ctrl}
	mock.recorder = &MockplatformifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockplatformifier) EXPECT() *MockplatformifierMockRecorder {
	return m.recorder
}

// Platformify mocks base method.
func (m *Mockplatformifier) Platformify(ctx context.Context, input *UserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Platformify", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// Platformify indicates an expected call of Platformify.
func (mr *MockplatformifierMockRecorder) Platformify(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Platformify", reflect.TypeOf((*Mockplatformifier)(nil).Platformify), ctx, input)
}

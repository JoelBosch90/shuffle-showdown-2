// Code generated by MockGen. DO NOT EDIT.
// Source: setup.go
//
// Generated by this command:
//
//	mockgen -source=setup.go -destination=mocks/App.go -package=mocks App
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	awscdk "github.com/aws/aws-cdk-go/awscdk/v2"
	cxapi "github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
	constructs "github.com/aws/constructs-go/constructs/v10"
	gomock "go.uber.org/mock/gomock"
)

// MockApp is a mock of App interface.
type MockApp struct {
	ctrl     *gomock.Controller
	recorder *MockAppMockRecorder
	isgomock struct{}
}

// MockAppMockRecorder is the mock recorder for MockApp.
type MockAppMockRecorder struct {
	mock *MockApp
}

// NewMockApp creates a new mock instance.
func NewMockApp(ctrl *gomock.Controller) *MockApp {
	mock := &MockApp{ctrl: ctrl}
	mock.recorder = &MockAppMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApp) EXPECT() *MockAppMockRecorder {
	return m.recorder
}

// Node mocks base method.
func (m *MockApp) Node() constructs.Node {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Node")
	ret0, _ := ret[0].(constructs.Node)
	return ret0
}

// Node indicates an expected call of Node.
func (mr *MockAppMockRecorder) Node() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*MockApp)(nil).Node))
}

// Synth mocks base method.
func (m *MockApp) Synth(options *awscdk.StageSynthesisOptions) cxapi.CloudAssembly {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Synth", options)
	ret0, _ := ret[0].(cxapi.CloudAssembly)
	return ret0
}

// Synth indicates an expected call of Synth.
func (mr *MockAppMockRecorder) Synth(options any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Synth", reflect.TypeOf((*MockApp)(nil).Synth), options)
}

// ToString mocks base method.
func (m *MockApp) ToString() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToString")
	ret0, _ := ret[0].(*string)
	return ret0
}

// ToString indicates an expected call of ToString.
func (mr *MockAppMockRecorder) ToString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToString", reflect.TypeOf((*MockApp)(nil).ToString))
}

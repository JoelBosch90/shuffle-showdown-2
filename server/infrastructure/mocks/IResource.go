// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces/IResource.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	awscdk "github.com/aws/aws-cdk-go/awscdk/v2"
	awsapigateway "github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	constructs "github.com/aws/constructs-go/constructs/v10"
	gomock "github.com/golang/mock/gomock"
)

// MockIResource is a mock of IResource interface.
type MockIResource struct {
	ctrl     *gomock.Controller
	recorder *MockIResourceMockRecorder
}

// MockIResourceMockRecorder is the mock recorder for MockIResource.
type MockIResourceMockRecorder struct {
	mock *MockIResource
}

// NewMockIResource creates a new mock instance.
func NewMockIResource(ctrl *gomock.Controller) *MockIResource {
	mock := &MockIResource{ctrl: ctrl}
	mock.recorder = &MockIResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIResource) EXPECT() *MockIResourceMockRecorder {
	return m.recorder
}

// AddCorsPreflight mocks base method.
func (m *MockIResource) AddCorsPreflight(options *awsapigateway.CorsOptions) awsapigateway.Method {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCorsPreflight", options)
	ret0, _ := ret[0].(awsapigateway.Method)
	return ret0
}

// AddCorsPreflight indicates an expected call of AddCorsPreflight.
func (mr *MockIResourceMockRecorder) AddCorsPreflight(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCorsPreflight", reflect.TypeOf((*MockIResource)(nil).AddCorsPreflight), options)
}

// AddMethod mocks base method.
func (m *MockIResource) AddMethod(httpMethod *string, target awsapigateway.Integration, options *awsapigateway.MethodOptions) awsapigateway.Method {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMethod", httpMethod, target, options)
	ret0, _ := ret[0].(awsapigateway.Method)
	return ret0
}

// AddMethod indicates an expected call of AddMethod.
func (mr *MockIResourceMockRecorder) AddMethod(httpMethod, target, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMethod", reflect.TypeOf((*MockIResource)(nil).AddMethod), httpMethod, target, options)
}

// AddProxy mocks base method.
func (m *MockIResource) AddProxy(options *awsapigateway.ProxyResourceOptions) awsapigateway.ProxyResource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProxy", options)
	ret0, _ := ret[0].(awsapigateway.ProxyResource)
	return ret0
}

// AddProxy indicates an expected call of AddProxy.
func (mr *MockIResourceMockRecorder) AddProxy(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProxy", reflect.TypeOf((*MockIResource)(nil).AddProxy), options)
}

// AddResource mocks base method.
func (m *MockIResource) AddResource(pathPart *string, options *awsapigateway.ResourceOptions) awsapigateway.Resource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddResource", pathPart, options)
	ret0, _ := ret[0].(awsapigateway.Resource)
	return ret0
}

// AddResource indicates an expected call of AddResource.
func (mr *MockIResourceMockRecorder) AddResource(pathPart, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResource", reflect.TypeOf((*MockIResource)(nil).AddResource), pathPart, options)
}

// Api mocks base method.
func (m *MockIResource) Api() awsapigateway.IRestApi {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Api")
	ret0, _ := ret[0].(awsapigateway.IRestApi)
	return ret0
}

// Api indicates an expected call of Api.
func (mr *MockIResourceMockRecorder) Api() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Api", reflect.TypeOf((*MockIResource)(nil).Api))
}

// ApplyRemovalPolicy mocks base method.
func (m *MockIResource) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ApplyRemovalPolicy", policy)
}

// ApplyRemovalPolicy indicates an expected call of ApplyRemovalPolicy.
func (mr *MockIResourceMockRecorder) ApplyRemovalPolicy(policy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyRemovalPolicy", reflect.TypeOf((*MockIResource)(nil).ApplyRemovalPolicy), policy)
}

// DefaultCorsPreflightOptions mocks base method.
func (m *MockIResource) DefaultCorsPreflightOptions() *awsapigateway.CorsOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultCorsPreflightOptions")
	ret0, _ := ret[0].(*awsapigateway.CorsOptions)
	return ret0
}

// DefaultCorsPreflightOptions indicates an expected call of DefaultCorsPreflightOptions.
func (mr *MockIResourceMockRecorder) DefaultCorsPreflightOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultCorsPreflightOptions", reflect.TypeOf((*MockIResource)(nil).DefaultCorsPreflightOptions))
}

// DefaultIntegration mocks base method.
func (m *MockIResource) DefaultIntegration() awsapigateway.Integration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultIntegration")
	ret0, _ := ret[0].(awsapigateway.Integration)
	return ret0
}

// DefaultIntegration indicates an expected call of DefaultIntegration.
func (mr *MockIResourceMockRecorder) DefaultIntegration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultIntegration", reflect.TypeOf((*MockIResource)(nil).DefaultIntegration))
}

// DefaultMethodOptions mocks base method.
func (m *MockIResource) DefaultMethodOptions() *awsapigateway.MethodOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultMethodOptions")
	ret0, _ := ret[0].(*awsapigateway.MethodOptions)
	return ret0
}

// DefaultMethodOptions indicates an expected call of DefaultMethodOptions.
func (mr *MockIResourceMockRecorder) DefaultMethodOptions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultMethodOptions", reflect.TypeOf((*MockIResource)(nil).DefaultMethodOptions))
}

// Env mocks base method.
func (m *MockIResource) Env() *awscdk.ResourceEnvironment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Env")
	ret0, _ := ret[0].(*awscdk.ResourceEnvironment)
	return ret0
}

// Env indicates an expected call of Env.
func (mr *MockIResourceMockRecorder) Env() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Env", reflect.TypeOf((*MockIResource)(nil).Env))
}

// GetResource mocks base method.
func (m *MockIResource) GetResource(pathPart *string) awsapigateway.IResource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", pathPart)
	ret0, _ := ret[0].(awsapigateway.IResource)
	return ret0
}

// GetResource indicates an expected call of GetResource.
func (mr *MockIResourceMockRecorder) GetResource(pathPart interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockIResource)(nil).GetResource), pathPart)
}

// Node mocks base method.
func (m *MockIResource) Node() constructs.Node {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Node")
	ret0, _ := ret[0].(constructs.Node)
	return ret0
}

// Node indicates an expected call of Node.
func (mr *MockIResourceMockRecorder) Node() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*MockIResource)(nil).Node))
}

// ParentResource mocks base method.
func (m *MockIResource) ParentResource() awsapigateway.IResource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParentResource")
	ret0, _ := ret[0].(awsapigateway.IResource)
	return ret0
}

// ParentResource indicates an expected call of ParentResource.
func (mr *MockIResourceMockRecorder) ParentResource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParentResource", reflect.TypeOf((*MockIResource)(nil).ParentResource))
}

// Path mocks base method.
func (m *MockIResource) Path() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(*string)
	return ret0
}

// Path indicates an expected call of Path.
func (mr *MockIResourceMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockIResource)(nil).Path))
}

// ResourceForPath mocks base method.
func (m *MockIResource) ResourceForPath(path *string) awsapigateway.Resource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceForPath", path)
	ret0, _ := ret[0].(awsapigateway.Resource)
	return ret0
}

// ResourceForPath indicates an expected call of ResourceForPath.
func (mr *MockIResourceMockRecorder) ResourceForPath(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceForPath", reflect.TypeOf((*MockIResource)(nil).ResourceForPath), path)
}

// ResourceId mocks base method.
func (m *MockIResource) ResourceId() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceId")
	ret0, _ := ret[0].(*string)
	return ret0
}

// ResourceId indicates an expected call of ResourceId.
func (mr *MockIResourceMockRecorder) ResourceId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceId", reflect.TypeOf((*MockIResource)(nil).ResourceId))
}

// Stack mocks base method.
func (m *MockIResource) Stack() awscdk.Stack {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stack")
	ret0, _ := ret[0].(awscdk.Stack)
	return ret0
}

// Stack indicates an expected call of Stack.
func (mr *MockIResourceMockRecorder) Stack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stack", reflect.TypeOf((*MockIResource)(nil).Stack))
}

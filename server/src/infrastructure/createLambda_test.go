package main

import (
	"infrastructure/interfaces"
	"infrastructure/mocks"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/golang/mock/gomock"
)

type MockNewFunction struct {
	mocks.Function
}

func (m *MockNewFunction) Get() interfaces.NewFunction {
	return func(scope constructs.Construct, id *string, props *awslambda.FunctionProps) awslambda.Function {
		m.SetTimesCalled(m.TimesCalled() + 1)

		return nil
	}
}

type MockNewLambdaRestApi struct {
	mocks.Function
	api interfaces.RestApi
}

func (m *MockNewLambdaRestApi) SetApi(api interfaces.RestApi) {
	m.api = api
}

func (m *MockNewLambdaRestApi) Get() interfaces.NewLambdaRestApi {
	return func(scope constructs.Construct, id *string, props *awsapigateway.LambdaRestApiProps) interfaces.RestApi {
		m.SetTimesCalled(m.TimesCalled() + 1)

		return m.api
	}
}

type MockNewLambdaIntegration struct {
	mocks.Function
}

func (m *MockNewLambdaIntegration) Get() interfaces.NewLambdaIntegration {
	return func(handler awslambda.IFunction, options *awsapigateway.LambdaIntegrationOptions) awsapigateway.LambdaIntegration {
		m.SetTimesCalled(m.TimesCalled() + 1)

		return nil
	}
}

func TestCreateLambda(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("interfaces.createLambda", func(t *testing.T) {
		// SETUP
		t.Parallel()

		// GIVEN
		mockStack := mocks.NewMockStack(controller)
		mockIResource := mocks.NewMockIResource(controller)
		mockResource := mocks.NewMockResource(controller)
		mockApi := mocks.NewMockRestApi(controller)
		mockApi.EXPECT().Root().Return(mockIResource).Times(1)
		mockIResource.EXPECT().AddResource(gomock.Any(), gomock.Any()).Return(mockResource).Times(1)
		mockResource.EXPECT().AddMethod(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)

		mockNewFunction := MockNewFunction{}
		mockNewApi := MockNewLambdaRestApi{}
		mockNewApi.SetApi(mockApi)
		mockNewIntegration := MockNewLambdaIntegration{}
		mockLambdaParameters := interfaces.LambdaParameters{
			Name:       "GreetFunction",
			SourcePath: "../controllers/greet",
			UrlPath:    "hello",
			Gateway:    "HelloWorldGateway",
		}

		// WHEN
		createLambda(mockStack, mockLambdaParameters, mockNewFunction.Get(), mockNewApi.Get(), mockNewIntegration.Get())

		// THEN
		newFunctionTimesCalled := mockNewFunction.TimesCalled()
		if newFunctionTimesCalled != 1 {
			t.Errorf("Expected newFunction to be called once, but was called %d times", newFunctionTimesCalled)
		}

		newApiTimesCalled := mockNewApi.TimesCalled()
		if newApiTimesCalled != 1 {
			t.Errorf("Expected newApi to be called once, but was called %d times", newApiTimesCalled)
		}

		newIntegrationTimesCalled := mockNewIntegration.TimesCalled()
		if newIntegrationTimesCalled != 1 {
			t.Errorf("Expected newIntegration to be called once, but was called %d times", newIntegrationTimesCalled)
		}
	})
}

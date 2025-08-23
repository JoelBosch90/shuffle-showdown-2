package main

import (
	"infrastructure/interfaces"
	"infrastructure/mocks"
	"os"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
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

type MockNewCfnOutput struct {
	mocks.Function
}

func (m *MockNewCfnOutput) Get() interfaces.NewCfnOutput {
	return func(scope constructs.Construct, id *string, props *awscdk.CfnOutputProps) awscdk.CfnOutput {
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
		mockUrl := "https://example.com"
		os.Setenv("PUBLIC_SHUFFLE_SHOWDOWN_API_PATH", "api")
		os.Setenv("PRIVATE_API_GATEWAY_URL_NAME", "ApiGatewayUrl")

		// GIVEN
		mockStack := mocks.NewMockStack(controller)
		mockApiProxyIResource := mocks.NewMockIResource(controller)
		mockLambdaResource := mocks.NewMockResource(controller)
		mockResource := mocks.NewMockResource(controller)
		mockApi := mocks.NewMockRestApi(controller)
		mockApi.EXPECT().Root().Return(mockApiProxyIResource).Times(1)
		mockApi.EXPECT().Url().Return(&mockUrl).Times(1)
		mockApiProxyIResource.EXPECT().AddResource(gomock.Any(), gomock.Any()).Return(mockResource).Times(1)
		mockResource.EXPECT().AddResource(gomock.Any(), gomock.Any()).Return(mockLambdaResource).Times(1)
		mockLambdaResource.EXPECT().AddMethod(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)

		mockNewFunction := MockNewFunction{}
		mockNewApi := MockNewLambdaRestApi{}
		mockNewApi.SetApi(mockApi)
		mockNewIntegration := MockNewLambdaIntegration{}
		mockNewCfnOutput := MockNewCfnOutput{}
		mockLambdaParameters := interfaces.LambdaParameters{
			Name:       "GreetFunction",
			SourcePath: "../controllers/greet",
			UrlPath:    "hello",
			Gateway:    "HelloWorldGateway",
		}

		// WHEN
		createLambda(mockStack, mockLambdaParameters, mockNewFunction.Get(), mockNewApi.Get(), mockNewIntegration.Get(), mockNewCfnOutput.Get())

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

		newCfnOutputTimesCalled := mockNewCfnOutput.TimesCalled()
		if newCfnOutputTimesCalled != 1 {
			t.Errorf("Expected newCfnOutput to be called once, but was called %d times", newCfnOutputTimesCalled)
		}
	})
}

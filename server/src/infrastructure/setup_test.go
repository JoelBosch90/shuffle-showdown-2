package main

import (
	"testing"

	"infrastructure/interfaces"
	"infrastructure/mocks"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/golang/mock/gomock"
)

type mockNewApp struct {
	mocks.Function
	app interfaces.App
}

func (m *mockNewApp) SetApp(app interfaces.App) {
	m.app = app
}

func (m *mockNewApp) Get() interfaces.NewApp {
	return func(props *awscdk.AppProps) interfaces.App {
		m.SetTimesCalled(m.TimesCalled() + 1)

		return m.app
	}
}

type mockNewStack struct {
	mocks.Function
}

func (m *mockNewStack) Get() interfaces.NewStack {
	return func(app interfaces.App, id *string, props *awscdk.StackProps) interfaces.Stack {
		m.SetTimesCalled(m.TimesCalled() + 1)

		return nil
	}
}

type mockCreateLambda struct {
	mocks.Function
}

func (m *mockCreateLambda) Get() interfaces.CreateLambda {
	return func(stack awscdk.Stack, params interfaces.LambdaParameters, newFunction interfaces.NewFunction, newApi interfaces.NewLambdaRestApi, newIntegration interfaces.NewLambdaIntegration) awslambda.Function {
		m.SetTimesCalled(m.TimesCalled() + 1)

		return nil
	}
}

type mockNewLambdaRestApi struct {
	mocks.Function
}

func (m *mockNewLambdaRestApi) Get() interfaces.NewLambdaRestApi {
	return func(scope constructs.Construct, id *string, props *awsapigateway.LambdaRestApiProps) interfaces.RestApi {
		m.SetTimesCalled(m.TimesCalled() + 1)

		return nil
	}
}

type mockNewLambdaIntegration struct {
	mocks.Function
}

func (m *mockNewLambdaIntegration) Get() interfaces.NewLambdaIntegration {
	return func(handler awslambda.IFunction, options *awsapigateway.LambdaIntegrationOptions) awsapigateway.LambdaIntegration {
		m.SetTimesCalled(m.TimesCalled() + 1)

		return nil
	}
}

func TestSetup(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("infrastructure.setup", func(t *testing.T) {
		// SETUP
		t.Parallel()

		// GIVEN
		mockApp := mocks.NewMockApp(controller)
		mockApp.EXPECT().Synth(nil).Times(1)
		mockNewApp := mockNewApp{}
		mockNewApp.SetApp(mockApp)
		mockNewStack := mockNewStack{}
		mockCreateLambda := mockCreateLambda{}
		mockNewLambdaRestApi := mockNewLambdaRestApi{}
		mockNewLambdaIntegration := mockNewLambdaIntegration{}
		mockCloseRuntime := mocks.Function{}

		// WHEN
		setup(mockNewApp.Get(), mockNewStack.Get(), mockCreateLambda.Get(), mockNewLambdaRestApi.Get(), mockNewLambdaIntegration.Get(), mockCloseRuntime.Get(), nil)

		// THEN
		appCreatorCalled := mockNewApp.TimesCalled()
		if appCreatorCalled != 1 {
			t.Errorf("Expected NewApp to be called once, but was called %d times", appCreatorCalled)
		}
		createStackCalled := mockNewStack.TimesCalled()
		if createStackCalled != 1 {
			t.Errorf("Expected NewStack to be called once, but was called %d times", createStackCalled)
		}
		createLambdaCalled := mockCreateLambda.TimesCalled()
		if createLambdaCalled != 1 {
			t.Errorf("Expected CreateLambda to be called once, but was called %d times", createLambdaCalled)
		}
		closeRuntimeCalled := mockCloseRuntime.TimesCalled()
		if closeRuntimeCalled != 1 {
			t.Errorf("Expected closeRuntime to be called once, but was called %d times", closeRuntimeCalled)
		}
	})
}

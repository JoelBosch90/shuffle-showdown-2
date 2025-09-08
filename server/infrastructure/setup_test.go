package main

import (
	"testing"

	"infrastructure/interfaces"
	"infrastructure/mocks"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
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

type mockCreateRestLambda struct {
	mocks.Function
}

func (m *mockCreateRestLambda) Get() interfaces.CreateRestLambda {
	return func(stack awscdk.Stack, params interfaces.RestLambdaParameters) awslambda.Function {
		m.SetTimesCalled(m.TimesCalled() + 1)

		return nil
	}
}

type mockCreateSocketLambdas struct {
	mocks.Function
}

func (m *mockCreateSocketLambdas) Get() interfaces.CreateSocketLambdas {
	return func(stack awscdk.Stack, params []interfaces.SocketLambdaParameters) awslambda.Function {
		m.SetTimesCalled(m.TimesCalled() + 1)

		return nil
	}
}

type mockCreateTable struct {
	mocks.Function
}

func (m *mockCreateTable) Get() interfaces.CreateTable {
	return func(stack awscdk.Stack, params interfaces.TableParameters) awsdynamodb.Table {
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
		mockCreateRestLambda := mockCreateRestLambda{}
		mockCreateSocketLambdas := mockCreateSocketLambdas{}
		mockCreateTable := mockCreateTable{}
		mockCloseRuntime := mocks.Function{}

		// WHEN
		setupWithDependencies(mockNewApp.Get(), mockNewStack.Get(), mockCreateRestLambda.Get(), mockCreateSocketLambdas.Get(), mockCreateTable.Get(), mockCloseRuntime.Get())

		// THEN
		appCreatorCalled := mockNewApp.TimesCalled()
		if appCreatorCalled != 1 {
			t.Errorf("Expected NewApp to be called once, but was called %d times", appCreatorCalled)
		}
		createStackCalled := mockNewStack.TimesCalled()
		if createStackCalled != 1 {
			t.Errorf("Expected NewStack to be called once, but was called %d times", createStackCalled)
		}
		createRestLambdaCalled := mockCreateRestLambda.TimesCalled()
		if createRestLambdaCalled != 1 {
			t.Errorf("Expected CreateRestLambda to be called once, but was called %d times", createRestLambdaCalled)
		}
		createSocketLambdaCalled := mockCreateSocketLambdas.TimesCalled()
		if createSocketLambdaCalled != 1 {
			t.Errorf("Expected CreateSocketLambda to be called once, but was called %d times", createSocketLambdaCalled)
		}
		createTableCalled := mockCreateTable.TimesCalled()
		if createTableCalled != 2 {
			t.Errorf("Expected CreateTable to be called twice, but was called %d times", createTableCalled)
		}
		closeRuntimeCalled := mockCloseRuntime.TimesCalled()
		if closeRuntimeCalled != 1 {
			t.Errorf("Expected closeRuntime to be called once, but was called %d times", closeRuntimeCalled)
		}
	})
}

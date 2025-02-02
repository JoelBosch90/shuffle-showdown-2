package main

import (
	"testing"

	"infrastructure/mocks"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func TestSetup(t *testing.T) {
	t.Run("infrastructure.setup", func(t *testing.T) {
		// SETUP
		t.Parallel()

		// GIVEN
		mockAppCreator := mocks.NewApp{}
		mockStackCreator := mocks.NewStack{}
		mockLambdaCreator := mocks.CreateLambda{}
		mockCloseRuntime := mocks.Function{}

		// WHEN
		setup(mockAppCreator.Get(), mockStackCreator.Get(), mockLambdaCreator.Get(), mockCloseRuntime.Get(), &awscdk.Environment{})

		// THEN
		appCreatorCalled := mockAppCreator.TimesCalled()
		if appCreatorCalled != 1 {
			t.Errorf("Expected createApp to be called once, but was called %d times", appCreatorCalled)
		}
		createStackCalled := mockStackCreator.TimesCalled()
		if createStackCalled != 1 {
			t.Errorf("Expected createStack to be called once, but was called %d times", createStackCalled)
		}
		createLambdaCalled := mockLambdaCreator.TimesCalled()
		if createLambdaCalled != 1 {
			t.Errorf("Expected createLambda to be called once, but was called %d times", createLambdaCalled)
		}
		closeRuntimeCalled := mockCloseRuntime.TimesCalled()
		if closeRuntimeCalled != 1 {
			t.Errorf("Expected closeRuntime to be called once, but was called %d times", closeRuntimeCalled)
		}
	})
}

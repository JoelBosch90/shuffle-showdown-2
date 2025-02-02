package main

import (
	"testing"
)

func TestCreateLambda(t *testing.T) {
	t.Run("infrastructure.lambda", func(t *testing.T) {
		// SETUP
		t.Parallel()

		// // GIVEN
		// mockAppCreator := mocks.AppCreator{}
		// mockStackCreator := mocks.StackCreator{}
		// mockLambdaCreator := mocks.LambdaCreator{}
		// mockCloseRuntime := mocks.Function{}

		// // WHEN
		// setup(mockAppCreator.GetFunction(), mockStackCreator.GetFunction(), mockLambdaCreator.GetFunction(), mockCloseRuntime.GetFunction(), &awscdk.Environment{})

		// // THEN
		// appCreatorCalled := mockAppCreator.GetTimesCalled()
		// if appCreatorCalled != 1 {
		// 	t.Errorf("Expected createApp to be called once, but was called %d times", appCreatorCalled)
		// }
		// createStackCalled := mockStackCreator.GetTimesCalled()
		// if createStackCalled != 1 {
		// 	t.Errorf("Expected createStack to be called once, but was called %d times", createStackCalled)
		// }
		// createLambdaCalled := mockLambdaCreator.GetTimesCalled()
		// if createLambdaCalled != 1 {
		// 	t.Errorf("Expected createLambda to be called once, but was called %d times", createLambdaCalled)
		// }
		// closeRuntimeCalled := mockCloseRuntime.GetTimesCalled()
		// if closeRuntimeCalled != 1 {
		// 	t.Errorf("Expected closeRuntime to be called once, but was called %d times", closeRuntimeCalled)
		// }
	})
}

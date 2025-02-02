package main

import (
	"testing"

	"infrastructure/mocks"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"go.uber.org/mock/gomock"
)

func TestSetup(t *testing.T) {
	t.Run("infrastructure.setup", func(t *testing.T) {
		// SETUP
		t.Parallel()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// GIVEN
		mockAppCreator := mocks.AppCreator{}
		mockStackCreator := mocks.Function{}
		mockLambdaCreator := mocks.LambdaCreator{}
		mockCloseRuntime := mocks.Function{}

		// WHEN
		setup(mockAppCreator.GetFunction(), mockStackCreator.GetFunction(), mockLambdaCreator.GetFunction(), mockCloseRuntime.GetFunction(), &awscdk.Environment{})

		// THEN
		// if closeRuntimeCalled != 1 {
		// 	t.Errorf("Expected mockCloseRuntime to be called once, but was called %d times", closeRuntimeCalled)
		// }
		// if appCreatorCalled != 1 {
		// 	t.Errorf("Expected mockAppCreator to be called once, but was called %d times", appCreatorCalled)
		// }
	})
}

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
		mockApp := mocks.NewMockAwsApp(ctrl)
		mockLambdaCollection := mocks.NewMockICollection(ctrl)

		closeRuntimeCalled := 0
		mockCloseRuntime := func() {
			closeRuntimeCalled++
		}

		appCreatorCalled := 0
		mockAppCreator := func(props *awscdk.AppProps) App {
			appCreatorCalled++
			return mockApp
		}

		// WHEN
		setup(mockCloseRuntime, mockAppCreator, nil, mockLambdaCollection)

		// THEN
		if closeRuntimeCalled != 1 {
			t.Errorf("Expected mockCloseRuntime to be called once, but was called %d times", closeRuntimeCalled)
		}
		if appCreatorCalled != 1 {
			t.Errorf("Expected mockAppCreator to be called once, but was called %d times", appCreatorCalled)
		}
		mockApp.EXPECT().Synth(gomock.Any()).Times(1)
		mockLambdaCollection.EXPECT().Create(gomock.Any()).Times(1)
	})
}

package main

import (
	"infrastructure/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreateLambda(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockStack := mocks.NewMockStack(controller)

	// // GIVEN
	// type testCase struct {
	// 	stack       awscdk.Stack
	// 	parameters  interfaces.LambdaParameters
	// 	newFunction interfaces.NewFunction
	// 	newApi      interfaces.NewLambdaRestApi
	// }

	// testCases := []testCase{
	// 	{
	// 		stack: mockStack,
	// 		parameters: interfaces.LambdaParameters{
	// 			Name:       "GreetFunction",
	// 			SourcePath: "../controllers/greet",
	// 			UrlPath:    "hello",
	// 			Gateway:    "HelloWorldGateway",
	// 		},
	// 		newFunction: mocks.NewFunction,
	// 		newApi:      mocks.NewLambdaRestApi,
	// 	},
	// }

	// for _, testCase := range testCases {
	// 	t.Run(testCase.name, func(t *testing.T) {
	// 		// SETUP
	// 		t.Parallel()

	// 		// GIVEN
	// 		mockAppCreator := mocks.AppCreator{}
	// 		mockStackCreator := mocks.StackCreator{}
	// 		mockLambdaCreator := mocks.LambdaCreator{}
	// 		mockCloseRuntime := mocks.Function{}

	// 		// WHEN
	// 		createLambda(mockAppCreator.GetFunction(), mockStackCreator.GetFunction(), mockLambdaCreator.GetFunction(), mockCloseRuntime.GetFunction(), &awscdk.Environment{})

	// 		// // THEN
	// 		// appCreatorCalled := mockAppCreator.GetTimesCalled()
	// 		// if appCreatorCalled != 1 {
	// 		// 	t.Errorf("Expected createApp to be called once, but was called %d times", appCreatorCalled)
	// 		// }
	// 		// createStackCalled := mockStackCreator.GetTimesCalled()
	// 		// if createStackCalled != 1 {
	// 		// 	t.Errorf("Expected createStack to be called once, but was called %d times", createStackCalled)
	// 		// }
	// 		// createLambdaCalled := mockLambdaCreator.GetTimesCalled()
	// 		// if createLambdaCalled != 1 {
	// 		// 	t.Errorf("Expected createLambda to be called once, but was called %d times", createLambdaCalled)
	// 		// }
	// 		// closeRuntimeCalled := mockCloseRuntime.GetTimesCalled()
	// 		// if closeRuntimeCalled != 1 {
	// 		// 	t.Errorf("Expected closeRuntime to be called once, but was called %d times", closeRuntimeCalled)
	// 		// }
	// 	})
	// }
}

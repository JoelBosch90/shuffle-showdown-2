package main

import (
	"infrastructure/interfaces"
	"infrastructure/mocks"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/golang/mock/gomock"
)

type MockNewWebSocketApi struct {
	mocks.Function
	api interfaces.WebSocketApi
}

func (m *MockNewWebSocketApi) SetApi(api interfaces.WebSocketApi) {
	m.api = api
}

func (m *MockNewWebSocketApi) Get() interfaces.NewWebSocketApi {
	return func(scope constructs.Construct, id *string, props *awsapigatewayv2.WebSocketApiProps) interfaces.WebSocketApi {
		m.SetTimesCalled(m.TimesCalled() + 1)
		return m.api
	}
}

type MockNewSocketFunction struct {
	mocks.Function
	functions []awslambda.Function
}

func (m *MockNewSocketFunction) SetFunctions(functions []awslambda.Function) {
	m.functions = functions
}

func (m *MockNewSocketFunction) Get() interfaces.NewFunction {
	return func(scope constructs.Construct, id *string, props *awslambda.FunctionProps) awslambda.Function {
		callIndex := m.TimesCalled()
		m.SetTimesCalled(m.TimesCalled() + 1)

		if callIndex < len(m.functions) {
			return m.functions[callIndex]
		}
		return nil
	}
}

type MockNewWebSocketLambdaIntegration struct {
	mocks.Function
}

func (m *MockNewWebSocketLambdaIntegration) Get() interfaces.NewWebSocketLambdaIntegration {
	return func(id *string, handler awslambda.IFunction, props *awsapigatewayv2integrations.WebSocketLambdaIntegrationProps) awsapigatewayv2integrations.WebSocketLambdaIntegration {
		m.SetTimesCalled(m.TimesCalled() + 1)

		return nil
	}
}

func TestCreateSocketLambdasWithDependencies(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	testCases := []struct {
		name                     string
		parameters               []interfaces.SocketLambdaParameters
		expectedLambdaCount      int
		expectedApiCalls         int
		expectedIntegrationCalls int
		expectedRouteCalls       int
	}{
		{
			name: "creates single WebSocket lambda with connect route",
			parameters: []interfaces.SocketLambdaParameters{
				{
					Name:       "ConnectFunction",
					SourcePath: "../controllers/websocket/connect",
					Route:      "$connect",
				},
			},
			expectedLambdaCount:      1,
			expectedApiCalls:         1,
			expectedIntegrationCalls: 1,
			expectedRouteCalls:       1,
		},
		{
			name: "creates multiple WebSocket lambdas with different routes",
			parameters: []interfaces.SocketLambdaParameters{
				{
					Name:       "ConnectFunction",
					SourcePath: "../controllers/websocket/connect",
					Route:      "$connect",
				},
				{
					Name:       "DisconnectFunction",
					SourcePath: "../controllers/websocket/disconnect",
					Route:      "$disconnect",
				},
				{
					Name:       "MessageFunction",
					SourcePath: "../controllers/websocket/message",
					Route:      "sendMessage",
				},
			},
			expectedLambdaCount:      3,
			expectedApiCalls:         1,
			expectedIntegrationCalls: 3,
			expectedRouteCalls:       3,
		},
		{
			name:                     "handles empty parameters array",
			parameters:               []interfaces.SocketLambdaParameters{},
			expectedLambdaCount:      0,
			expectedApiCalls:         1,
			expectedIntegrationCalls: 0,
			expectedRouteCalls:       0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// SETUP
			mockStack := mocks.NewMockStack(controller)

			// Setup mock tables for each parameter
			var updatedParams []interfaces.SocketLambdaParameters
			for _, param := range testCase.parameters {
				mockTable := mocks.NewMockTable(controller)
				mockTable.EXPECT().GrantReadWriteData(gomock.Any()).Times(1)
				mockTable.EXPECT().TableName().Return(jsii.String("test-table")).Times(1)

				updatedParam := param
				updatedParam.Table = mockTable
				updatedParams = append(updatedParams, updatedParam)
			}

			// Setup mock WebSocket API
			mockWebSocketApi := mocks.NewMockWebSocketApi(controller)
			mockWebSocketApi.EXPECT().ApiEndpoint().Return(jsii.String("wss://test-api.execute-api.region.amazonaws.com/stage")).Times(len(testCase.parameters))
			mockWebSocketApi.EXPECT().AddRoute(gomock.Any(), gomock.Any()).Times(testCase.expectedRouteCalls)

			mockNewWebSocketApi := MockNewWebSocketApi{}
			mockNewWebSocketApi.SetApi(mockWebSocketApi)

			// Setup mock Lambda functions
			var mockFunctions []awslambda.Function
			for i := 0; i < testCase.expectedLambdaCount; i++ {
				mockFunction := mocks.NewMockFunction(controller)
				mockFunctions = append(mockFunctions, mockFunction)
			}

			mockNewFunction := MockNewSocketFunction{}
			mockNewFunction.SetFunctions(mockFunctions)

			mockNewIntegration := MockNewWebSocketLambdaIntegration{}

			// WHEN
			result := createSocketLambdaWithDependencies(
				mockStack,
				updatedParams,
				mockNewWebSocketApi.Get(),
				mockNewFunction.Get(),
				mockNewIntegration.Get(),
			)

			// THEN
			if len(result) != testCase.expectedLambdaCount {
				t.Errorf("Expected %d lambdas to be returned, but got %d", testCase.expectedLambdaCount, len(result))
			}

			// Verify WebSocket API was created once
			if mockNewWebSocketApi.TimesCalled() != testCase.expectedApiCalls {
				t.Errorf("Expected WebSocket API to be created %d times, but was called %d times",
					testCase.expectedApiCalls, mockNewWebSocketApi.TimesCalled())
			}

			// Verify Lambda functions were created
			if mockNewFunction.TimesCalled() != testCase.expectedLambdaCount {
				t.Errorf("Expected %d Lambda functions to be created, but %d were created",
					testCase.expectedLambdaCount, mockNewFunction.TimesCalled())
			}

			// Verify integrations were created
			if mockNewIntegration.TimesCalled() != testCase.expectedIntegrationCalls {
				t.Errorf("Expected %d integrations to be created, but %d were created",
					testCase.expectedIntegrationCalls, mockNewIntegration.TimesCalled())
			}

			// Verify returned lambdas match expected functions
			for i, lambda := range result {
				if lambda != mockFunctions[i] {
					t.Errorf("Returned lambda at index %d does not match expected mock function", i)
				}
			}
		})
	}
}

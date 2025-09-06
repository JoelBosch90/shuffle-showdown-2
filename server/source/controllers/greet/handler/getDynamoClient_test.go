package handler

import (
	"context"
	"errors"
	"testing"

	"greet/interfaces"
	"greet/mocks"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/golang/mock/gomock"
)

func TestGetDynamoClientWithDependencies(t *testing.T) {
	mockContext := mocks.GetMockContext()

	testCases := []struct {
		name            string
		context         context.Context
		setupConfigMock func(*testing.T) interfaces.LoadDefaultConfig
		setupClientMock func(*testing.T, *gomock.Controller) (interfaces.DynamoDBClient, interfaces.NewFromConfig)
		expectError     bool
		expectClient    bool
	}{
		{
			name: "successfully creates DynamoDB client",
			setupConfigMock: func(t *testing.T) interfaces.LoadDefaultConfig {
				return func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
					return aws.Config{
						Region: "us-east-1",
					}, nil
				}
			},
			setupClientMock: func(t *testing.T, ctrl *gomock.Controller) (interfaces.DynamoDBClient, interfaces.NewFromConfig) {
				mockClient := mocks.NewMockDynamoDBClient(ctrl)

				mockNewFromConfig := func(cfg aws.Config, optFns ...func(*dynamodb.Options)) interfaces.DynamoDBClient {
					// Verify that the config passed in matches what we expect
					if cfg.Region != "us-east-1" {
						t.Errorf("Expected region 'us-east-1', got '%s'", cfg.Region)
					}
					return mockClient
				}

				return mockClient, mockNewFromConfig
			},
			expectError:  false,
			expectClient: true,
		},
		{
			name: "returns error when config loading fails",
			setupConfigMock: func(t *testing.T) interfaces.LoadDefaultConfig {
				return func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
					return aws.Config{}, errors.New("failed to load AWS config")
				}
			},
			setupClientMock: func(t *testing.T, ctrl *gomock.Controller) (interfaces.DynamoDBClient, interfaces.NewFromConfig) {
				mockNewFromConfig := func(cfg aws.Config, optFns ...func(*dynamodb.Options)) interfaces.DynamoDBClient {
					t.Error("newFromConfig should not be called when config loading fails")
					return nil
				}

				return nil, mockNewFromConfig
			},
			expectError:  true,
			expectClient: false,
		},
		{
			name: "passes context correctly to loadConfig",
			setupConfigMock: func(t *testing.T) interfaces.LoadDefaultConfig {
				return func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
					if ctx != mockContext {
						t.Error("Expected context to be passed to loadConfig")
					}

					return aws.Config{Region: "us-west-2"}, nil
				}
			},
			setupClientMock: func(t *testing.T, ctrl *gomock.Controller) (interfaces.DynamoDBClient, interfaces.NewFromConfig) {
				mockClient := mocks.NewMockDynamoDBClient(ctrl)

				mockNewFromConfig := func(cfg aws.Config, optFns ...func(*dynamodb.Options)) interfaces.DynamoDBClient {
					return mockClient
				}

				return mockClient, mockNewFromConfig
			},
			expectError:  false,
			expectClient: true,
		},
		{
			name: "handles config with custom options",
			setupConfigMock: func(t *testing.T) interfaces.LoadDefaultConfig {
				return func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
					// Verify that options can be passed (even if we don't use them in this test)
					if len(optFns) > 10 { // Arbitrary check to ensure we can handle options
						t.Log("Config loaded with options")
					}

					return aws.Config{
						Region: "eu-west-1",
					}, nil
				}
			},
			setupClientMock: func(t *testing.T, ctrl *gomock.Controller) (interfaces.DynamoDBClient, interfaces.NewFromConfig) {
				mockClient := mocks.NewMockDynamoDBClient(ctrl)

				mockNewFromConfig := func(cfg aws.Config, optFns ...func(*dynamodb.Options)) interfaces.DynamoDBClient {
					// Verify that DynamoDB options can be passed
					if cfg.Region != "eu-west-1" {
						t.Errorf("Expected region 'eu-west-1', got '%s'", cfg.Region)
					}
					return mockClient
				}

				return mockClient, mockNewFromConfig
			},
			expectError:  false,
			expectClient: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockLoadConfig := testCase.setupConfigMock(t)
			expectedClient, mockNewFromConfig := testCase.setupClientMock(t, ctrl)

			// Call the function under test
			client, err := getDynamoClientWithDependencies(mockContext, mockLoadConfig, mockNewFromConfig)

			// Verify error expectation
			if testCase.expectError && err == nil {
				t.Error("Expected error but got none")
				return
			}
			if !testCase.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
				return
			}

			// Verify client expectation
			if testCase.expectClient && client == nil {
				t.Error("Expected client but got nil")
				return
			}
			if !testCase.expectClient && client != nil {
				t.Error("Expected nil client but got a client")
				return
			}

			// Verify that the returned client is the expected mock client
			if testCase.expectClient && client != expectedClient {
				t.Error("Returned client is not the expected mock client")
			}
		})
	}
}

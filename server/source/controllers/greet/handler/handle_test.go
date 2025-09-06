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
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/mock/gomock"
)

func TestHandle(t *testing.T) {
	testCases := []struct {
		name                 string
		setupConfigMock      func() interfaces.LoadDefaultConfig
		setupGetMessageMock  func() func(context.Context) (*dynamodb.GetItemOutput, error)
		setupSetMessageMock  func() func(context.Context) error
		expectedStatusCode   int
		expectedBodyContains string
		expectError          bool
	}{
		{
			name: "returns existing message when found on first getMessage call",
			setupConfigMock: func() interfaces.LoadDefaultConfig {
				return func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
					return aws.Config{}, nil
				}
			},
			setupGetMessageMock: func() func(context.Context) (*dynamodb.GetItemOutput, error) {
				callCount := 0

				return func(ctx context.Context) (*dynamodb.GetItemOutput, error) {
					callCount++

					if callCount > 1 {
						t.Error("getMessage should only be called once when message exists")
					}

					return &dynamodb.GetItemOutput{
						Item: map[string]types.AttributeValue{
							MessagePartitionPropertyName: &types.AttributeValueMemberS{Value: "Existing message from database"},
						},
					}, nil
				}
			},
			setupSetMessageMock: func() func(context.Context) error {
				return func(ctx context.Context) error {
					t.Error("setMessage should not be called when message exists")

					return nil
				}
			},
			expectedStatusCode:   200,
			expectedBodyContains: "Existing message from database",
			expectError:          false,
		},
		{
			name: "creates new message when not found, then returns it",
			setupConfigMock: func() interfaces.LoadDefaultConfig {
				return func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
					return aws.Config{}, nil
				}
			},
			setupGetMessageMock: func() func(context.Context) (*dynamodb.GetItemOutput, error) {
				callCount := 0

				return func(ctx context.Context) (*dynamodb.GetItemOutput, error) {
					callCount++

					switch callCount {
					case 1:
						// First call returns empty
						return &dynamodb.GetItemOutput{Item: nil}, errors.New("Mocked database error")
					case 2:
						// Second call returns the created message
						return &dynamodb.GetItemOutput{
							Item: map[string]types.AttributeValue{
								MessagePartitionPropertyName: &types.AttributeValueMemberS{Value: "Hello from the database, lovely world!"},
							},
						}, nil
					default:
						t.Errorf("getMessage called too many times: %d", callCount)
						return nil, errors.New("too many calls")
					}
				}
			},
			setupSetMessageMock: func() func(context.Context) error {
				callCount := 0

				return func(ctx context.Context) error {
					callCount++

					if callCount > 1 {
						t.Error("setMessage should only be called once")
					}

					return nil
				}
			},
			expectedStatusCode:   200,
			expectedBodyContains: "Hello from the database, lovely world!",
			expectError:          false,
		},
		{
			name: "returns 500 when setMessage fails",
			setupConfigMock: func() interfaces.LoadDefaultConfig {
				return func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
					return aws.Config{}, nil
				}
			},
			setupGetMessageMock: func() func(context.Context) (*dynamodb.GetItemOutput, error) {
				callCount := 0

				return func(ctx context.Context) (*dynamodb.GetItemOutput, error) {
					callCount++

					if callCount > 1 {
						t.Error("getMessage should only be called once when setMessage fails")
					}
					return &dynamodb.GetItemOutput{Item: nil}, errors.New("Mocked database error")
				}
			},
			setupSetMessageMock: func() func(context.Context) error {
				return func(ctx context.Context) error {
					return &types.ResourceNotFoundException{}
				}
			},
			expectedStatusCode:   500,
			expectedBodyContains: "Failed to set database record",
			expectError:          true,
		},
		{
			name: "returns 500 when getMessage fails twice",
			setupConfigMock: func() interfaces.LoadDefaultConfig {
				return func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
					return aws.Config{}, nil
				}
			},
			setupGetMessageMock: func() func(context.Context) (*dynamodb.GetItemOutput, error) {
				callCount := 0

				return func(ctx context.Context) (*dynamodb.GetItemOutput, error) {
					callCount++

					if callCount > 2 {
						t.Error("getMessage should only be called twice when it fails twice")
					}
					return &dynamodb.GetItemOutput{Item: nil}, errors.New("Mocked database error")
				}
			},
			setupSetMessageMock: func() func(context.Context) error {
				callCount := 0

				return func(ctx context.Context) error {
					callCount++

					if callCount > 1 {
						t.Error("setMessage should only be called once")
					}

					return nil
				}
			},
			expectedStatusCode:   500,
			expectedBodyContains: "Failed to get database record",
			expectError:          true,
		},
		{
			name: "returns 500 when getMessage returns an invalid item",
			setupConfigMock: func() interfaces.LoadDefaultConfig {
				return func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (aws.Config, error) {
					return aws.Config{}, nil
				}
			},
			setupGetMessageMock: func() func(context.Context) (*dynamodb.GetItemOutput, error) {
				callCount := 0

				return func(ctx context.Context) (*dynamodb.GetItemOutput, error) {
					callCount++

					if callCount > 1 {
						t.Error("getMessage should only be called once when message exists")
					}

					return &dynamodb.GetItemOutput{
						Item: map[string]types.AttributeValue{
							"SomeOtherAttribute": &types.AttributeValueMemberS{Value: "some value"},
						},
					}, nil
				}
			},
			setupSetMessageMock: func() func(context.Context) error {
				callCount := 0

				return func(ctx context.Context) error {
					callCount++

					if callCount > 1 {
						t.Error("setMessage should only be called once")
					}

					return nil
				}
			},
			expectedStatusCode:   500,
			expectedBodyContains: "Failed to retrieve message",
			expectError:          true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockContext := mocks.GetMockContext()
			mockGetMessage := testCase.setupGetMessageMock()
			mockSetMessage := testCase.setupSetMessageMock()

			response, err := HandleWithDependencies(mockContext, mockGetMessage, mockSetMessage)

			if testCase.expectError && err == nil {
				t.Error("Expected error but got none")
				return
			}
			if !testCase.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
				return
			}

			if response.StatusCode != testCase.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", testCase.expectedStatusCode, response.StatusCode)
			}

			if testCase.expectedBodyContains != "" && response.Body != testCase.expectedBodyContains {
				t.Errorf("Expected body to contain '%s', got '%s'", testCase.expectedBodyContains, response.Body)
			}
		})
	}
}

package handler

import (
	"context"
	"errors"
	"testing"

	"greet/interfaces"
	"greet/mocks"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
			name: "creates new message when getMessage returns empty item",
			setupGetMessageMock: func() func(context.Context) (*dynamodb.GetItemOutput, error) {
				callCount := 0

				return func(ctx context.Context) (*dynamodb.GetItemOutput, error) {
					callCount++

					switch callCount {
					case 1:
						// First call returns empty item (no error - this is normal DynamoDB behavior)
						return &dynamodb.GetItemOutput{Item: map[string]types.AttributeValue{}}, nil
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
			name: "creates new message when getMessage returns nil item",
			setupGetMessageMock: func() func(context.Context) (*dynamodb.GetItemOutput, error) {
				callCount := 0

				return func(ctx context.Context) (*dynamodb.GetItemOutput, error) {
					callCount++

					switch callCount {
					case 1:
						// First call returns nil item (no error - normal DynamoDB behavior for missing item)
						return &dynamodb.GetItemOutput{Item: nil}, nil
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
			name: "creates new message when item exists but missing required attribute",
			setupGetMessageMock: func() func(context.Context) (*dynamodb.GetItemOutput, error) {
				callCount := 0

				return func(ctx context.Context) (*dynamodb.GetItemOutput, error) {
					callCount++

					switch callCount {
					case 1:
						// First call returns item without required attribute
						return &dynamodb.GetItemOutput{
							Item: map[string]types.AttributeValue{
								"SomeOtherAttribute": &types.AttributeValueMemberS{Value: "some value"},
								// Missing MessagePartitionPropertyName
							},
						}, nil
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
			name: "returns 500 when first getMessage fails with error",
			setupGetMessageMock: func() func(context.Context) (*dynamodb.GetItemOutput, error) {
				return func(ctx context.Context) (*dynamodb.GetItemOutput, error) {
					return nil, errors.New("DynamoDB connection error")
				}
			},
			setupSetMessageMock: func() func(context.Context) error {
				return func(ctx context.Context) error {
					t.Error("setMessage should not be called when first getMessage fails")
					return nil
				}
			},
			expectedStatusCode:   500,
			expectedBodyContains: "Failed to retrieve message",
			expectError:          true,
		},
		{
			name: "returns 500 when setMessage fails",
			setupGetMessageMock: func() func(context.Context) (*dynamodb.GetItemOutput, error) {
				return func(ctx context.Context) (*dynamodb.GetItemOutput, error) {
					// Return empty item to trigger setMessage call
					return &dynamodb.GetItemOutput{Item: nil}, nil
				}
			},
			setupSetMessageMock: func() func(context.Context) error {
				return func(ctx context.Context) error {
					return errors.New("Failed to create item")
				}
			},
			expectedStatusCode:   500,
			expectedBodyContains: "Failed to save message",
			expectError:          true,
		},
		{
			name: "returns 500 when second getMessage fails after setMessage succeeds",
			setupGetMessageMock: func() func(context.Context) (*dynamodb.GetItemOutput, error) {
				callCount := 0

				return func(ctx context.Context) (*dynamodb.GetItemOutput, error) {
					callCount++

					switch callCount {
					case 1:
						// First call returns empty (triggers setMessage)
						return &dynamodb.GetItemOutput{Item: nil}, nil
					case 2:
						// Second call fails
						return nil, errors.New("Database error on second call")
					default:
						t.Errorf("getMessage called too many times: %d", callCount)
						return nil, errors.New("too many calls")
					}
				}
			},
			setupSetMessageMock: func() func(context.Context) error {
				return func(ctx context.Context) error {
					return nil // setMessage succeeds
				}
			},
			expectedStatusCode:   500,
			expectedBodyContains: "Failed to retrieve message",
			expectError:          true,
		},
		{
			name: "returns 500 when item has wrong attribute type",
			setupGetMessageMock: func() func(context.Context) (*dynamodb.GetItemOutput, error) {
				callCount := 0

				return func(ctx context.Context) (*dynamodb.GetItemOutput, error) {
					callCount++

					switch callCount {
					case 1:
						// First call returns item with wrong attribute type
						return &dynamodb.GetItemOutput{
							Item: map[string]types.AttributeValue{
								MessagePartitionPropertyName: &types.AttributeValueMemberN{Value: "123"}, // Number instead of string
							},
						}, nil
					case 2:
						// Second call returns item with wrong attribute type again
						return &dynamodb.GetItemOutput{
							Item: map[string]types.AttributeValue{
								MessagePartitionPropertyName: &types.AttributeValueMemberN{Value: "123"}, // Number instead of string
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
			expectedStatusCode:   500,
			expectedBodyContains: "Failed to retrieve message",
			expectError:          true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
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

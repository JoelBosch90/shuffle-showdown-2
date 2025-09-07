package handler

import (
	"context"
	"os"
	"testing"

	"greet/mocks"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/mock/gomock"
)

var TestMockMessage = &dynamodb.GetItemOutput{
	Item: map[string]types.AttributeValue{
		MessagePartitionPropertyName: &types.AttributeValueMemberS{Value: "Hello from the database, lovely world!"},
	},
}

func TestGetMessage(t *testing.T) {
	testCases := []struct {
		name        string
		setupMock   func(*mocks.MockDynamoDBClient, context.Context)
		expectError bool
	}{
		{
			name: "calls GetItem with correct parameters",
			setupMock: func(mockClient *mocks.MockDynamoDBClient, expectedCtx context.Context) {
				// Create the expected input
				expectedInput := &dynamodb.GetItemInput{
					TableName: aws.String(os.Getenv("TABLE_NAME")),
					Key: map[string]types.AttributeValue{
						MessagePartitionKeyName: &types.AttributeValueMemberS{Value: MessagePartitionKeyValue},
					},
				}

				mockClient.EXPECT().
					GetItem(gomock.Eq(expectedCtx), gomock.Any(), gomock.Any()).
					Return(TestMockMessage, nil).
					Times(1).
					Do(func(ctx context.Context, input *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) {
						if *input.TableName != *expectedInput.TableName {
							t.Errorf("Expected TableName %s, got %s", *expectedInput.TableName, *input.TableName)
						}

						if input.Key == nil {
							t.Error("Expected Key to be set, but it was nil")
							return
						}

						value, ok := input.Key[MessagePartitionKeyName]
						if !ok {
							t.Errorf("Expected Key to contain %s, but it was missing", MessagePartitionKeyName)
							return
						}

						attribute, ok := value.(*types.AttributeValueMemberS)
						if !ok {
							t.Errorf("Expected Key[%s] to be of type *AttributeValueMemberS, but got %T", MessagePartitionKeyName, value)
							return
						}

						if attribute.Value != MessagePartitionKeyValue {
							t.Errorf("Expected Key[%s] to have value %s, but got %s", MessagePartitionKeyName, MessagePartitionKeyValue, attribute.Value)
						}
					})
			},
			expectError: false,
		},
		{
			name: "returns error when DynamoDB fails",
			setupMock: func(mockClient *mocks.MockDynamoDBClient, expectedCtx context.Context) {
				mockClient.EXPECT().
					GetItem(gomock.Eq(expectedCtx), gomock.Any(), gomock.Any()).
					Return(nil, &types.ResourceNotFoundException{}).
					Times(1)
			},
			expectError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			os.Setenv("TABLE_NAME", "HelloWorldTable")

			mockContext := mocks.GetMockContext()
			mockClient := mocks.NewMockDynamoDBClient(ctrl)

			testCase.setupMock(mockClient, mockContext)

			messageOutput, error := getMessageHandleWithDependencies(mockContext, mockClient)

			if testCase.expectError && error == nil {
				t.Error("Expected error but got none")
				return
			}
			if testCase.expectError && error != nil {
				// Test passed, we expected an error and got one
				return
			}

			if error != nil {
				t.Errorf("Expected no error but got: %v", error)
				return
			}

			if messageOutput == nil {
				t.Error("Expected message but got nil")
				return
			}

			if messageOutput.Item == nil {
				t.Error("Expected message.Item to be set but got nil")
				return
			}

			messageAttribute, messageAttributeExists := messageOutput.Item[MessagePartitionPropertyName]
			if !messageAttributeExists {
				t.Errorf("Expected message.Item to contain key %s but it was missing", MessagePartitionPropertyName)
				return
			}

			message, messageOk := messageAttribute.(*types.AttributeValueMemberS)
			if !messageOk {
				t.Errorf("Expected message.Item[%s] to be of type *AttributeValueMemberS but got %T", MessagePartitionPropertyName, messageAttribute)
				return
			}

			if message.Value != "Hello from the database, lovely world!" {
				t.Errorf("Expected message.Item[%s] to have value 'Hello from the database, lovely world!' but got %s", MessagePartitionPropertyName, message.Value)
			}
		})
	}
}

package handler

import (
	"context"
	"testing"

	"greet/mocks"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang/mock/gomock"
)

func TestSetMessage(t *testing.T) {
	testCases := []struct {
		name        string
		setupMock   func(*mocks.MockDynamoDBClient, context.Context)
		expectError bool
	}{
		{
			name: "calls PutItem with correct parameters",
			setupMock: func(mockClient *mocks.MockDynamoDBClient, expectedCtx context.Context) {
				// Create the expected input
				expectedInput := &dynamodb.PutItemInput{
					TableName: aws.String(TableName),
					Item: map[string]types.AttributeValue{
						MessagePartitionKeyName:      &types.AttributeValueMemberS{Value: MessagePartitionKeyValue},
						MessagePartitionPropertyName: &types.AttributeValueMemberS{Value: "Hello from the database, lovely world!"},
					},
				}

				mockClient.EXPECT().
					PutItem(gomock.Eq(expectedCtx), gomock.Any(), gomock.Any()).
					Return(&dynamodb.PutItemOutput{}, nil).
					Times(1).
					Do(func(ctx context.Context, input *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) {
						if *input.TableName != *expectedInput.TableName {
							t.Errorf("Expected TableName %s, got %s", *expectedInput.TableName, *input.TableName)
						}

						if input.Item == nil {
							t.Error("Expected Item to be set, but it was nil")
							return
						}

						if len(input.Item) != len(expectedInput.Item) {
							t.Errorf("Expected Item to have %d entries, got %d", len(expectedInput.Item), len(input.Item))
						}

						partitionKeyValue, partitionKeyExists := input.Item[MessagePartitionKeyName]
						if !partitionKeyExists {
							t.Errorf("Expected Item to contain key %s, but it was missing", MessagePartitionKeyName)
							return
						}

						partitionKeyAttribute, partitionKeyOk := partitionKeyValue.(*types.AttributeValueMemberS)
						if !partitionKeyOk {
							t.Errorf("Expected Item[%s] to be of type *AttributeValueMemberS, but got %T", MessagePartitionKeyName, partitionKeyValue)
							return
						}

						if partitionKeyAttribute.Value != MessagePartitionKeyValue {
							t.Errorf("Expected Item[%s] to have value %s, but got %s", MessagePartitionKeyName, MessagePartitionKeyValue, partitionKeyAttribute.Value)
						}

						messageValue, messageExists := input.Item[MessagePartitionPropertyName]
						if !messageExists {
							t.Errorf("Expected Item to contain key %s, but it was missing", MessagePartitionPropertyName)
							return
						}

						messageAttribute, messageOk := messageValue.(*types.AttributeValueMemberS)
						if !messageOk {
							t.Errorf("Expected Item[%s] to be of type *AttributeValueMemberS, but got %T", MessagePartitionPropertyName, messageValue)
							return
						}

						if messageAttribute.Value != "Hello from the database, lovely world!" {
							t.Errorf("Expected Item[%s] to have value 'Hello from the database, lovely world!', but got %s", MessagePartitionPropertyName, messageAttribute.Value)
						}
					})
			},
			expectError: false,
		},
		{
			name: "returns error when DynamoDB fails",
			setupMock: func(mockClient *mocks.MockDynamoDBClient, expectedCtx context.Context) {
				mockClient.EXPECT().
					PutItem(gomock.Eq(expectedCtx), gomock.Any(), gomock.Any()).
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

			mockContext := mocks.GetMockContext()
			mockClient := mocks.NewMockDynamoDBClient(ctrl)

			testCase.setupMock(mockClient, mockContext)

			error := setMessageHandleWithDependencies(mockContext, mockClient)

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
		})
	}
}

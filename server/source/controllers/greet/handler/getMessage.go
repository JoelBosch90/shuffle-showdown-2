package handler

import (
	"context"
	"greet/interfaces"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func getMessageHandleWithDependencies(ctx context.Context, dynamoClient interfaces.DynamoDBClient) (*dynamodb.GetItemOutput, error) {
	return dynamoClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Key: map[string]types.AttributeValue{
			MessagePartitionKeyName: &types.AttributeValueMemberS{Value: MessagePartitionKeyValue},
		},
	})
}

package handler

import (
	"context"
	"greet/interfaces"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func setMessageHandleWithDependencies(ctx context.Context, dynamoClient interfaces.DynamoDBClient) error {
	_, error := dynamoClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Item: map[string]types.AttributeValue{
			MessagePartitionKeyName:      &types.AttributeValueMemberS{Value: MessagePartitionKeyValue},
			MessagePartitionPropertyName: &types.AttributeValueMemberS{Value: "Hello from the database, lovely world!"},
		},
	})

	return error
}

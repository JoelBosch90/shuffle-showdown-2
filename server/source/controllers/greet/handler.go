package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Handler func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

const tableName = "HelloWorldTable"
const partitionKeyName = "PK"
const partitionKeyValue = "greeting"
const partitionPropertyName = "Message"

func setDatabaseRecord(ctx context.Context, dynamoClient *dynamodb.Client) error {
	_, err := dynamoClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			partitionKeyName:      &types.AttributeValueMemberS{Value: partitionKeyValue},
			partitionPropertyName: &types.AttributeValueMemberS{Value: "Hello from the database, lovely world!"},
		},
	})
	return err
}

func getDatabaseRecord(ctx context.Context, dynamoClient *dynamodb.Client) (*dynamodb.GetItemOutput, error) {
	return dynamoClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			partitionKeyName: &types.AttributeValueMemberS{Value: partitionKeyValue},
		},
	})
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg, configError := config.LoadDefaultConfig(ctx)
	if configError != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to load config",
		}, configError
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)

	result, getError := getDatabaseRecord(ctx, dynamoClient)
	if getError == nil || result.Item == nil {
		setError := setDatabaseRecord(ctx, dynamoClient)
		if setError != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Failed to set database record",
			}, setError
		}

		result, getError = getDatabaseRecord(ctx, dynamoClient)
		if getError != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Failed to get database record",
			}, getError
		}
	}

	if messageAttr, exists := result.Item[partitionPropertyName]; exists {
		if message, ok := messageAttr.(*types.AttributeValueMemberS); ok {
			return events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       message.Value,
			}, nil
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       "Failed to retrieve message",
	}, nil
}

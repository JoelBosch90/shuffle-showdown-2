package handler

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func HandleWithDependencies(
	ctx context.Context,
	callGetMessage func(context.Context) (*dynamodb.GetItemOutput, error),
	callSetMessage func(context.Context) error,
) (events.APIGatewayProxyResponse, error) {
	result, getError := callGetMessage(ctx)
	if getError != nil {
		setError := callSetMessage(ctx)
		if setError != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Failed to set database record",
			}, setError
		}

		result, getError = callGetMessage(ctx)
		if getError != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Failed to get database record",
			}, getError
		}
	}

	log.Println("DynamoDB GetItem result:", result, getError)

	messageAttr, exists := result.Item[MessagePartitionPropertyName]
	if exists {
		message, ok := messageAttr.(*types.AttributeValueMemberS)
		if ok {
			return events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       message.Value,
			}, nil
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       "Failed to retrieve message",
	}, errors.New("message attribute missing or of wrong type")
}

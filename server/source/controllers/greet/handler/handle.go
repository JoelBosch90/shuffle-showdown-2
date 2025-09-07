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
	log.Printf("First getMessage - error: %v, item count: %d", getError, len(result.Item))

	if getError != nil {
		log.Println("Creating message...")
		setError := callSetMessage(ctx)
		log.Printf("setMessage - error: %v", setError)
		if setError != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Failed to set database record",
			}, setError
		}

		result, getError = callGetMessage(ctx)
		log.Printf("Second getMessage - error: %v, item count: %d", getError, len(result.Item))

		if getError != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Failed to get database record",
			}, getError
		}
	}

	log.Printf("Final result.Item: %+v", result.Item)
	log.Printf("Looking for key: %s", MessagePartitionPropertyName)

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

	log.Printf("Message attribute missing or of wrong type in result.Item: %+v with messageAttr %+v and error %v", result.Item, messageAttr, getError)

	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       "Failed to retrieve message",
	}, errors.New("message attribute missing or of wrong type")
}

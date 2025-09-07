package handler

import (
	"context"
	"errors"

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
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to retrieve message",
		}, getError
	}

	messageAttribute, exists := result.Item[MessagePartitionPropertyName]
	message, ok := messageAttribute.(*types.AttributeValueMemberS)
	if len(result.Item) > 0 && exists && ok {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       message.Value,
		}, nil
	}

	setError := callSetMessage(ctx)
	if setError != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to save message",
		}, setError
	}

	result2, getError2 := callGetMessage(ctx)
	if getError2 != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Failed to retrieve message",
		}, getError2
	}

	messageAttribute2, exists2 := result2.Item[MessagePartitionPropertyName]
	message2, ok2 := messageAttribute2.(*types.AttributeValueMemberS)
	if len(result2.Item) > 0 && exists2 && ok2 {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       message2.Value,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       "Failed to retrieve message",
	}, errors.New("message attribute missing or of wrong type")
}

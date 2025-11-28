package handler

import (
	"context"
	"errors"

	"github.com/aws/aws-lambda-go/events"
)

func HandleWithDependencies(
	ctx context.Context,
) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       "Failed to do anything",
	}, errors.New("bla bla bla")
}

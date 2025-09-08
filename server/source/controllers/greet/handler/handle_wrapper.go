package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

func Handle(
	ctx context.Context,
) (events.APIGatewayProxyResponse, error) {

	return HandleWithDependencies(ctx, getMessage, setMessage)
}

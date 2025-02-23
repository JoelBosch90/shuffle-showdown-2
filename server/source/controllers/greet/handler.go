package main

import (
	"github.com/aws/aws-lambda-go/events"
)

type Handler func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello, wordly world!",
	}, nil
}

package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	// GIVEN
	type testCase struct {
		name     string
		request  events.APIGatewayProxyRequest
		response events.APIGatewayProxyResponse
		error    error
	}

	testCases := []testCase{
		{
			name:    "greet.handler: Accepts an empty request",
			request: events.APIGatewayProxyRequest{},
			response: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "Hello, wordly world!",
			},
			error: nil,
		},
		{
			name: "greet.handler: Ignores the request body",
			request: events.APIGatewayProxyRequest{
				Body: "Hello, world!",
			},
			response: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "Hello, wordly world!",
			},
			error: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			// SETUP
			t.Parallel()

			// WHEN
			got, error := handler(testCase.request)

			// THEN
			if error != testCase.error {
				t.Errorf("handler returned unexpected error: %v", error)
			}

			if got.StatusCode != testCase.response.StatusCode || got.Body != testCase.response.Body {
				t.Errorf("handler = %v, want %v", got, testCase.response)
			}
		})
	}
}

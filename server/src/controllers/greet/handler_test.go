package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandler(t *testing.T) {
	// GIVEN
	var tests = []struct {
		name  string
		given events.APIGatewayProxyRequest
		want  events.APIGatewayProxyResponse
		error error
	}{
		{
			name:  "Accepts an empty request",
			given: events.APIGatewayProxyRequest{},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "Hello, wordy world!",
			},
			error: nil,
		},
		{
			name: "Ignores the request body",
			given: events.APIGatewayProxyRequest{
				Body: "Hello, world!",
			},
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "Hello, wordy world!",
			},
			error: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// SETUP
			t.Parallel()

			// WHEN
			got, error := handler(test.given)

			// THEN
			if error != test.error {
				t.Errorf("handler returned unexpected error: %v", error)
			}
			if got.StatusCode != test.want.StatusCode || got.Body != test.want.Body {
				t.Errorf("handler = %v, want %v", got, test.want)
			}
		})
	}
}

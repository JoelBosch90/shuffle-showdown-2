//go:build !skip_test
// +build !skip_test

package main

import (
	"greet/handler"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.Handle)
}

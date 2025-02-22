//go:build !skip_test
// +build !skip_test

package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	setup(lambda.Start)
}

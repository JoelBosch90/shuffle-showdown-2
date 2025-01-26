//go:build !skip_test
// +build !skip_test

package main

import (
	"infrastructure/lambda"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	setup(jsii.Close, func(props *awscdk.AppProps) App { return awscdk.NewApp(props) }, nil, &lambda.Collection{})
}

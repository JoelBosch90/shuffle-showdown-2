//go:build !skip_test
// +build !skip_test

package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

func main() {
	setup(
		func(props *awscdk.AppProps) interfaces.App {
			return awscdk.NewApp(props)
		},
		func(app interfaces.App, stackId *string, props *awscdk.StackProps) interfaces.Stack {
			return awscdk.NewStack(app, stackId, props)
		},
		func(stack awscdk.Stack, params interfaces.LambdaParameters, newFunction interfaces.NewFunction, newApi interfaces.NewLambdaRestApi) awslambda.Function {
			return createLambda(stack, params, newFunction, newApi)
		},
		jsii.Close,
		nil,
	)
}

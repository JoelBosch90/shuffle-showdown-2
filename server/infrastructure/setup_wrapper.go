package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
)

func setup() {
	createApp := func(props *awscdk.AppProps) interfaces.App {
		return awscdk.NewApp(props)
	}
	createStack := func(app interfaces.App, stackId *string, props *awscdk.StackProps) interfaces.Stack {
		return awscdk.NewStack(app, stackId, props)
	}

	setupWithDependencies(
		createApp,
		createStack,
		createRestLambda,
		createSocketLambdas,
		createTable,
		jsii.Close,
	)
}

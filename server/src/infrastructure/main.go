//go:build !skip_test
// +build !skip_test

package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/constructs-go/constructs/v10"
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
		createLambda,
		func(scope constructs.Construct, id *string, props *awsapigateway.LambdaRestApiProps) interfaces.RestApi {
			return awsapigateway.NewLambdaRestApi(scope, id, props)
		},
		awsapigateway.NewLambdaIntegration,
		jsii.Close,
		nil,
	)
}

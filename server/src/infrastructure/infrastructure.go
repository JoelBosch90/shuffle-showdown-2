//go:build !skip_test
// +build !skip_test

package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
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
		func(stack awscdk.Stack, params interfaces.LambdaParameters, newFunction interfaces.NewFunction, newApi interfaces.NewLambdaRestApi, newIntegration interfaces.NewLambdaIntegration) awslambda.Function {
			return createLambda(stack, params, newFunction, newApi, newIntegration)
		},
		func(scope constructs.Construct, id *string, props *awsapigateway.LambdaRestApiProps) interfaces.RestApi {
			return awsapigateway.NewLambdaRestApi(scope, id, props)
		},
		func(handler awslambda.IFunction, options *awsapigateway.LambdaIntegrationOptions) awsapigateway.LambdaIntegration {
			return awsapigateway.NewLambdaIntegration(handler, options)
		},
		jsii.Close,
		nil,
	)
}

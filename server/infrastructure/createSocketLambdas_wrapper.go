package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

func createSocketLambdas(stack awscdk.Stack, lambdaParams []interfaces.SocketLambdaParameters) []awslambda.Function {
	newWebSocketApi := func(scope constructs.Construct, id *string, props *awsapigatewayv2.WebSocketApiProps) interfaces.WebSocketApi {
		return awsapigatewayv2.NewWebSocketApi(scope, id, props)
	}

	return createSocketLambdaWithDependencies(stack, lambdaParams, newWebSocketApi, awslambda.NewFunction, awsapigatewayv2integrations.NewWebSocketLambdaIntegration)
}

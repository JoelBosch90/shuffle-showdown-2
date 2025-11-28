package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

func createSocketLambdaWithDependencies(stack awscdk.Stack, parameters []interfaces.SocketLambdaParameters, newWebSocketApi interfaces.NewWebSocketApi, newFunction interfaces.NewFunction, newWebSocketLambdaIntegration interfaces.NewWebSocketLambdaIntegration) []awslambda.Function {
	webSocketApi := newWebSocketApi(stack, jsii.String("WebSocketApi"), &awsapigatewayv2.WebSocketApiProps{
		ApiName: jsii.String("socket"),
	})

	var lambdas []awslambda.Function
	for _, param := range parameters {
		lambda := newFunction(stack, jsii.String(param.Name), &awslambda.FunctionProps{
			Runtime:      awslambda.Runtime_PROVIDED_AL2(),
			Handler:      jsii.String("bootstrap"),
			Code:         awslambda.Code_FromAsset(jsii.String(param.SourcePath), nil),
			Architecture: awslambda.Architecture_ARM_64(),
			Environment: &map[string]*string{
				"TABLE_NAME":    param.Table.TableName(),
				"WEBSOCKET_URL": webSocketApi.ApiEndpoint(),
			},
		})

		param.Table.GrantReadWriteData(lambda)

		integration := newWebSocketLambdaIntegration(
			jsii.String(param.Name+"Integration"),
			lambda,
			&awsapigatewayv2integrations.WebSocketLambdaIntegrationProps{},
		)

		webSocketApi.AddRoute(jsii.String(param.Route), &awsapigatewayv2.WebSocketRouteOptions{
			Integration: integration,
		})

		lambdas = append(lambdas, lambda)
	}

	return lambdas
}

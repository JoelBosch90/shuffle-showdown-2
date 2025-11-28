package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type NewWebSocketLambdaIntegration func(id *string, handler awslambda.IFunction, props *awsapigatewayv2integrations.WebSocketLambdaIntegrationProps) awsapigatewayv2integrations.WebSocketLambdaIntegration

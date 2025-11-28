package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/constructs-go/constructs/v10"
)

type NewWebSocketApi func(scope constructs.Construct, id *string, props *awsapigatewayv2.WebSocketApiProps) WebSocketApi

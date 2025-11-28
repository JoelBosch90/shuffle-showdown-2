package interfaces

import "github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"

type WebSocketApi interface {
	awsapigatewayv2.WebSocketApi
}

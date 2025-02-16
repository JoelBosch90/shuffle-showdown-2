package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type NewLambdaIntegration func(handler awslambda.IFunction, options *awsapigateway.LambdaIntegrationOptions) awsapigateway.LambdaIntegration

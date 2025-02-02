package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/constructs-go/constructs/v10"
)

type NewLambdaRestApi func(scope constructs.Construct, id *string, props *awsapigateway.LambdaRestApiProps) RestApi

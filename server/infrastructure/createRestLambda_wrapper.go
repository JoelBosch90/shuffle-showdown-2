package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
)

func createRestLambda(stack awscdk.Stack, parameters interfaces.RestLambdaParameters) awslambda.Function {
	newApi := func(scope constructs.Construct, id *string, props *awsapigateway.LambdaRestApiProps) interfaces.RestApi {
		return awsapigateway.NewLambdaRestApi(scope, id, props)
	}

	return createRestLambdaWithDependencies(stack, parameters, awslambda.NewFunction, newApi, awsapigateway.NewLambdaIntegration, awsssm.NewStringParameter)
}

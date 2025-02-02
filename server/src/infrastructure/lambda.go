package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

func CreateLambda(stack awscdk.Stack, parameters interfaces.LambdaParameters) awslambda.Function {
	lambda := awslambda.NewFunction(stack, jsii.String(parameters.Name), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String(parameters.SourcePath), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	gateway := awsapigateway.NewLambdaRestApi(stack, jsii.String(parameters.Gateway), &awsapigateway.LambdaRestApiProps{
		Handler: lambda,
		Proxy:   jsii.Bool(true),
	})

	helloResource := gateway.Root().AddResource(jsii.String(parameters.UrlPath), &awsapigateway.ResourceOptions{})
	helloResource.AddMethod(jsii.String("GET"), awsapigateway.NewLambdaIntegration(lambda, &awsapigateway.LambdaIntegrationOptions{}), &awsapigateway.MethodOptions{})

	return lambda
}

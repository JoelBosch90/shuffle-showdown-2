package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

func createLambda(stack awscdk.Stack, parameters interfaces.LambdaParameters, newFunction interfaces.NewFunction, newApi interfaces.NewLambdaRestApi) awslambda.Function {
	lambda := newFunction(stack, jsii.String(parameters.Name), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String(parameters.SourcePath), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	api := newApi(stack, jsii.String(parameters.Gateway), &awsapigateway.LambdaRestApiProps{
		Handler: lambda,
		Proxy:   jsii.Bool(true),
	})

	helloResource := api.Root().AddResource(jsii.String(parameters.UrlPath), &awsapigateway.ResourceOptions{})
	helloResource.AddMethod(jsii.String("GET"), awsapigateway.NewLambdaIntegration(lambda, &awsapigateway.LambdaIntegrationOptions{}), &awsapigateway.MethodOptions{})

	return lambda
}

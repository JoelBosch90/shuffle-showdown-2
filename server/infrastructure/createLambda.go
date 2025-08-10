package main

import (
	"infrastructure/interfaces"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

func createLambda(stack awscdk.Stack, parameters interfaces.LambdaParameters, newFunction interfaces.NewFunction, newApi interfaces.NewLambdaRestApi, newIntegration interfaces.NewLambdaIntegration, newCfnOutput interfaces.NewCfnOutput) awslambda.Function {
	lambda := newFunction(stack, jsii.String(parameters.Name), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String(parameters.SourcePath), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	api := newApi(stack, jsii.String(parameters.Gateway), &awsapigateway.LambdaRestApiProps{
		Handler: lambda,
		Proxy:   jsii.Bool(false),
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: jsii.Strings("https://" + os.Getenv("SHUFFLE_SHOWDOWN_DOMAIN")),
			AllowMethods: jsii.Strings("GET", "POST", "PUT", "DELETE", "OPTIONS"),
			AllowHeaders: jsii.Strings("Content-Type", "Authorization"),
			MaxAge:       awscdk.Duration_Seconds(jsii.Number(300)),
		},
	})

	apiProxy := api.Root().AddResource(jsii.String("api"), &awsapigateway.ResourceOptions{})
	integration := newIntegration(lambda, &awsapigateway.LambdaIntegrationOptions{})

	helloResource := apiProxy.AddResource(jsii.String(parameters.UrlPath), &awsapigateway.ResourceOptions{})
	helloResource.AddMethod(jsii.String("GET"), integration, &awsapigateway.MethodOptions{})

	newCfnOutput(stack, jsii.String(os.Getenv("API_GATEWAY_NAME")), &awscdk.CfnOutputProps{
		Value:       api.Url(),
		Description: jsii.String("The URL of the API Gateway"),
	})

	return lambda
}

package main

import (
	"infrastructure/interfaces"
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/jsii-runtime-go"
)

func createLambda(stack awscdk.Stack, parameters interfaces.LambdaParameters, newFunction interfaces.NewFunction, newApi interfaces.NewLambdaRestApi, newIntegration interfaces.NewLambdaIntegration, newStringParameter interfaces.NewStringParameter) awslambda.Function {
	lambda := newFunction(stack, jsii.String(parameters.Name), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String(parameters.SourcePath), nil),
		Architecture: awslambda.Architecture_ARM_64(),
		Environment: &map[string]*string{
			"TABLE_NAME": parameters.Table.TableName(),
		},
	})

	lambda.AddToRolePolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Effect:    awsiam.Effect_ALLOW,
		Actions:   jsii.Strings("dynamodb:GetItem", "dynamodb:PutItem"),
		Resources: jsii.Strings(*parameters.Table.TableArn()),
	}))

	api := newApi(stack, jsii.String(parameters.Gateway), &awsapigateway.LambdaRestApiProps{
		Handler: lambda,
		Proxy:   jsii.Bool(false),
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: jsii.Strings("https://" + os.Getenv("PUBLIC_SHUFFLE_SHOWDOWN_DOMAIN")),
			AllowMethods: jsii.Strings("GET", "POST", "PUT", "DELETE", "OPTIONS"),
			AllowHeaders: jsii.Strings("Content-Type", "Authorization", "Origin"),
			MaxAge:       awscdk.Duration_Seconds(jsii.Number(300)),
		},
		DeployOptions: &awsapigateway.StageOptions{
			StageName: jsii.String(os.Getenv("PUBLIC_STAGE")),
		},
	})

	apiProxy := api.Root().AddResource(jsii.String(os.Getenv("PUBLIC_SHUFFLE_SHOWDOWN_API_PATH")), &awsapigateway.ResourceOptions{})
	integration := newIntegration(lambda, &awsapigateway.LambdaIntegrationOptions{})

	lambdaEndpoint := apiProxy.AddResource(jsii.String(parameters.UrlPath), &awsapigateway.ResourceOptions{})
	lambdaEndpoint.AddMethod(jsii.String("GET"), integration, &awsapigateway.MethodOptions{})

	apiGateWayUrlName := os.Getenv("PRIVATE_API_GATEWAY_URL_NAME")
	newStringParameter(stack, jsii.String(apiGateWayUrlName), &awsssm.StringParameterProps{
		ParameterName: jsii.String(apiGateWayUrlName),
		StringValue:   api.Url(),
		Description:   jsii.String("API Gateway URL for the application"),
	})

	parameters.Table.GrantReadWriteData(lambda)
	lambda.AddEnvironment(jsii.String("TABLE_NAME"), parameters.Table.TableName(), nil)

	return lambda
}

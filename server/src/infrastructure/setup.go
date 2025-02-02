package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

func setup(createApp interfaces.NewApp, createStack interfaces.NewStack, createLambda interfaces.CreateLambda, closeRunTime func(), environment *awscdk.Environment) {
	defer closeRunTime()

	app := createApp(nil)

	stackId := "InfrastructureStack"
	stack := createStack(
		app,
		&stackId,
		&awscdk.StackProps{
			Env: environment,
		},
	)

	lambdasToCreate := []interfaces.LambdaParameters{
		{
			Name:       "GreetFunction",
			SourcePath: "../controllers/greet",
			UrlPath:    "hello",
			Gateway:    "HelloWorldGateway",
		},
	}

	for _, lambdaParams := range lambdasToCreate {
		createLambda(
			stack,
			lambdaParams,
			awslambda.NewFunction,
			func(scope constructs.Construct, id *string, props *awsapigateway.LambdaRestApiProps) interfaces.RestApi {
				return awsapigateway.NewLambdaRestApi(scope, id, props)
			},
		)
	}

	app.Synth(nil)
}

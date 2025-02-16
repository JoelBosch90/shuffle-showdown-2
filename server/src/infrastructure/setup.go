package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func setup(
	createApp interfaces.NewApp,
	createStack interfaces.NewStack,
	createLambda interfaces.CreateLambda,
	newLambdaRestApi interfaces.NewLambdaRestApi,
	newLambdaIntegration interfaces.NewLambdaIntegration,
	closeRunTime func(),
	environment *awscdk.Environment,
) {
	defer closeRunTime()

	app := createApp(nil)

	stackId := "ServerStack"
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
			newLambdaRestApi,
			newLambdaIntegration,
		)
	}

	app.Synth(nil)
}

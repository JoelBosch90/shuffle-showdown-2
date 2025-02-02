package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func setup(createApp interfaces.AppCreator, createStack interfaces.StackCreator, createLambda interfaces.LambdaCreator, closeRunTime func(), environment *awscdk.Environment) {
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
		createLambda(stack, lambdaParams)
	}

	app.Synth(nil)
}

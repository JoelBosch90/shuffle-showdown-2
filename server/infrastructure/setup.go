package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
)

func setup(
	createApp interfaces.NewApp,
	createStack interfaces.NewStack,
	createLambda interfaces.CreateLambda,
	createTable interfaces.CreateTable,
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

	helloWorldTable := createTable(stack, interfaces.TableParameters{
		ID:               "HelloWorldTable",
		PartitionKeyName: "PK",
	}, awsdynamodb.NewTable)

	lambdasToCreate := []interfaces.LambdaParameters{
		{
			Name:       "GreetFunction",
			SourcePath: "../source/controllers/greet",
			UrlPath:    "hello",
			Gateway:    "HelloWorldGateway",
			Table:      helloWorldTable,
		},
	}

	for _, lambdaParams := range lambdasToCreate {
		createLambda(
			stack,
			lambdaParams,
			awslambda.NewFunction,
			newLambdaRestApi,
			newLambdaIntegration,
			awsssm.NewStringParameter,
		)
	}

	app.Synth(nil)
}

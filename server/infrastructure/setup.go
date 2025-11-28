package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func setupWithDependencies(
	createAppFunc interfaces.NewApp,
	createStackFunc interfaces.NewStack,
	createRestLambdaFunc interfaces.CreateRestLambda,
	createSocketLambdasFunc interfaces.CreateSocketLambdas,
	createTableFunc interfaces.CreateTable,
	closeRunTime func(),
) {
	defer closeRunTime()

	app := createAppFunc(nil)

	stackId := "ServerStack"
	stack := createStackFunc(
		app,
		&stackId,
		&awscdk.StackProps{},
	)

	helloWorldTable := createTableFunc(stack, interfaces.TableParameters{
		ID:               "HelloWorldTable",
		PartitionKeyName: "PK",
	})

	restLambdasToCreate := []interfaces.RestLambdaParameters{
		{
			Name:       "GreetFunction",
			SourcePath: "../source/controllers/greet",
			Route:      "hello",
			Gateway:    "HelloWorldGateway",
			Table:      helloWorldTable,
		},
	}

	for _, restLambdaParams := range restLambdasToCreate {
		createRestLambdaFunc(stack, restLambdaParams)
	}

	connectionsTable := createTableFunc(stack, interfaces.TableParameters{
		ID:               "Connections",
		PartitionKeyName: "PK",
	})

	socketLambdasToCreate := []interfaces.SocketLambdaParameters{
		{
			Name:       "ConnectFunction",
			SourcePath: "../source/controllers/connect",
			Route:      "$connect",
			Table:      connectionsTable,
		},
	}

	createSocketLambdasFunc(stack, socketLambdasToCreate)

	app.Synth(nil)
}

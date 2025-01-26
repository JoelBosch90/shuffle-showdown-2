package main

import (
	"infrastructure/lambda"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
	"github.com/aws/constructs-go/constructs/v10"
)

type App interface {
	constructs.Construct
	Synth(options *awscdk.StageSynthesisOptions) cxapi.CloudAssembly
}

type AppCreator func(props *awscdk.AppProps) App

func setup(closeRunTime func(), createApp AppCreator, environment *awscdk.Environment, lambdaCollection lambda.ICollection) {
	defer closeRunTime()

	app := createApp(nil)

	stackId := "InfrastructureStack"
	stack := awscdk.NewStack(
		app,
		&stackId,
		&awscdk.StackProps{
			Env: environment,
		},
	)

	lambdaCollection.Create(stack)

	app.Synth(nil)
}

package mocks

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
	"github.com/aws/constructs-go/constructs/v10"
)

type App struct{}

func (a *App) Node() constructs.Node {
	return nil
}

func (a *App) ToString() *string {
	return nil
}

func (a *App) Synth(options *awscdk.StageSynthesisOptions) cxapi.CloudAssembly {
	return nil
}

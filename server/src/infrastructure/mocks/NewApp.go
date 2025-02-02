package mocks

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type NewApp struct {
	Function
}

func (m *NewApp) Get() interfaces.NewApp {
	return func(props *awscdk.AppProps) interfaces.App {
		m.timesCalled++

		return nil
	}
}

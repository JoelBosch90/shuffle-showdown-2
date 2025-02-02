package mocks

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type AppCreator struct {
	Function
}

func (m *AppCreator) GetFunction() interfaces.AppCreator {
	return func(props *awscdk.AppProps) interfaces.App {
		m.timesCalled += 1

		return &App{}
	}
}

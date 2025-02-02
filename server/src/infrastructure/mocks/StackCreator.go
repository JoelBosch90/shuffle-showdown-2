package mocks

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type StackCreator struct {
	Function
}

func (m *StackCreator) GetFunction() interfaces.StackCreator {
	return func(app interfaces.App, id *string, props *awscdk.StackProps) awscdk.Stack {
		m.timesCalled++

		return nil
	}
}

package mocks

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type NewStack struct {
	Function
}

func (m *NewStack) Get() interfaces.NewStack {
	return func(app interfaces.App, id *string, props *awscdk.StackProps) interfaces.Stack {
		m.timesCalled++

		return nil
	}
}

package mocks

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

type NewFunction struct {
	Function
}

func (m *NewFunction) Get() interfaces.NewFunction {
	return func(scope constructs.Construct, id *string, props *awslambda.FunctionProps) awslambda.Function {
		m.timesCalled++

		return nil
	}
}

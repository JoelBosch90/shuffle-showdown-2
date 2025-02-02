package mocks

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type LambdaCreator struct {
	Function
}

func (m *LambdaCreator) GetFunction() interfaces.LambdaCreator {
	return func(stack awscdk.Stack, params interfaces.LambdaParameters) awslambda.Function {
		m.timesCalled++

		return nil
	}
}

package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
)

type NewFunction func(scope constructs.Construct, id *string, props *awslambda.FunctionProps) awslambda.Function

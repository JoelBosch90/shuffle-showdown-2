package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type RestLambdaParameters struct {
	Name       string
	SourcePath string
	Route      string
	Gateway    string
	Table      awsdynamodb.Table
}

type CreateRestLambda func(stack awscdk.Stack, params RestLambdaParameters) awslambda.Function

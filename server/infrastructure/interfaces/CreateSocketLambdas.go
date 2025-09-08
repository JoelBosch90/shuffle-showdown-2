package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type SocketLambdaParameters struct {
	Name       string
	SourcePath string
	Route      string
	Table      awsdynamodb.Table
}

type CreateSocketLambdas func(stack awscdk.Stack, params []SocketLambdaParameters) awslambda.Function

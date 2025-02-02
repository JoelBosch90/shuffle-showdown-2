package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type LambdaParameters struct {
	Name       string
	SourcePath string
	UrlPath    string
	Gateway    string
}

type LambdaCreator func(stack awscdk.Stack, params LambdaParameters) awslambda.Function

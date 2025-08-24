package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

type LambdaParameters struct {
	Name       string
	SourcePath string
	UrlPath    string
	Gateway    string
	Table      awsdynamodb.Table
}

type CreateLambda func(stack awscdk.Stack, params LambdaParameters, newFunction NewFunction, newApi NewLambdaRestApi, newIntegration NewLambdaIntegration, newStringParameter NewStringParameter) awslambda.Function

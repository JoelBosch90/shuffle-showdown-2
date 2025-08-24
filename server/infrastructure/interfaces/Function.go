package interfaces

import "github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

type Function interface {
	awslambda.Function
}

package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm"
	"github.com/aws/constructs-go/constructs/v10"
)

type NewStringParameter func(scope constructs.Construct, id *string, props *awsssm.StringParameterProps) awsssm.StringParameter

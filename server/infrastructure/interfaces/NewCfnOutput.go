package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type NewCfnOutput func(scope constructs.Construct, id *string, props *awscdk.CfnOutputProps) awscdk.CfnOutput

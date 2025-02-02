package interfaces

import "github.com/aws/aws-cdk-go/awscdk/v2"

type NewStack func(app App, id *string, props *awscdk.StackProps) Stack

package interfaces

import "github.com/aws/aws-cdk-go/awscdk/v2"

type StackCreator func(app App, id *string, props *awscdk.StackProps) awscdk.Stack

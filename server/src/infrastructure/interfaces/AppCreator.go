package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type AppCreator func(props *awscdk.AppProps) App

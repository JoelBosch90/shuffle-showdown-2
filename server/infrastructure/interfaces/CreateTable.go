package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
)

type TableParameters struct {
	ID               string
	PartitionKeyName string
}

type CreateTable func(stack awscdk.Stack, params TableParameters) awsdynamodb.Table

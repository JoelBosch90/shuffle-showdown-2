package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
)

func createTable(stack awscdk.Stack, parameters interfaces.TableParameters) awsdynamodb.Table {
	return createTableWithDependencies(stack, parameters, awsdynamodb.NewTable)
}

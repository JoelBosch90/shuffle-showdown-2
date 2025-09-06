package main

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/jsii-runtime-go"
)

func createTable(stack awscdk.Stack, parameters interfaces.TableParameters, newTable interfaces.NewTable) awsdynamodb.Table {
	return newTable(stack, jsii.String(parameters.ID), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String(parameters.PartitionKeyName),
			Type: awsdynamodb.AttributeType_STRING,
		},
		BillingMode:   awsdynamodb.BillingMode_PAY_PER_REQUEST,
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})
}

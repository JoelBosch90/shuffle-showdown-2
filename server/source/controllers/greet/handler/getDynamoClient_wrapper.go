package handler

import (
	"context"
	"greet/interfaces"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func getDynamoClient(ctx context.Context) (interfaces.DynamoDBClient, error) {
	newFromConfigWrapper := func(cfg aws.Config, optFns ...func(*dynamodb.Options)) interfaces.DynamoDBClient {
		return dynamodb.NewFromConfig(cfg, optFns...)
	}

	return getDynamoClientWithDependencies(ctx, config.LoadDefaultConfig, newFromConfigWrapper)
}

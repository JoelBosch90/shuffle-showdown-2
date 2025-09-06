package handler

import (
	"context"
	"greet/interfaces"
)

func getDynamoClientWithDependencies(ctx context.Context, loadConfig interfaces.LoadDefaultConfig, newFromConfig interfaces.NewFromConfig) (interfaces.DynamoDBClient, error) {
	cfg, configError := loadConfig(ctx)
	if configError != nil {
		return nil, configError
	}

	dynamoClient := newFromConfig(cfg)

	return dynamoClient, nil
}

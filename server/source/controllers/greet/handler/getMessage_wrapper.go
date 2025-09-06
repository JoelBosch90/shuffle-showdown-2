//go:build !skip_test
// +build !skip_test

package handler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func getMessage(ctx context.Context) (*dynamodb.GetItemOutput, error) {
	client, error := getDynamoClient(ctx)
	if error != nil {
		return nil, error
	}

	return getMessageHandleWithDependencies(ctx, client)
}

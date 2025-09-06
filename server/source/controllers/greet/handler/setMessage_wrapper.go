//go:build !skip_test
// +build !skip_test

package handler

import (
	"context"
)

func setMessage(ctx context.Context) error {
	client, error := getDynamoClient(ctx)
	if error != nil {
		return error
	}

	return setMessageHandleWithDependencies(ctx, client)
}

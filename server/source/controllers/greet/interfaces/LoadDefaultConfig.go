package interfaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type LoadDefaultConfig func(ctx context.Context, optFns ...func(*config.LoadOptions) error) (cfg aws.Config, err error)

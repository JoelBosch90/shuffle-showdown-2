package interfaces

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type NewFromConfig func(cfg aws.Config, optFns ...func(*dynamodb.Options)) DynamoDBClient

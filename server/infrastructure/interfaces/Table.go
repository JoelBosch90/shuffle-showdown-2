package interfaces

import "github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"

type Table interface {
	awsdynamodb.Table
}

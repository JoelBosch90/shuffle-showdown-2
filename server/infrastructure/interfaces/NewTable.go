package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
)

type NewTable func(scope constructs.Construct, id *string, props *awsdynamodb.TableProps) awsdynamodb.Table

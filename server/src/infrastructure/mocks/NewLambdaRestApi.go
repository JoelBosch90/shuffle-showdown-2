package mocks

import (
	"infrastructure/interfaces"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/constructs-go/constructs/v10"
)

type NewLambdaRestApi struct {
	Function
}

func (m *NewLambdaRestApi) Get() interfaces.NewLambdaRestApi {
	return func(scope constructs.Construct, id *string, props *awsapigateway.LambdaRestApiProps) interfaces.RestApi {
		m.timesCalled++

		return nil
	}
}

package interfaces

import "github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"

type IResource interface {
	awsapigateway.IResource
}

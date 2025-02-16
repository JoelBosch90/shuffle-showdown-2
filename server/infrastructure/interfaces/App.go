package interfaces

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/cxapi"
	"github.com/aws/constructs-go/constructs/v10"
)

type App interface {
	Node() constructs.Node
	ToString() *string
	Synth(options *awscdk.StageSynthesisOptions) cxapi.CloudAssembly
}

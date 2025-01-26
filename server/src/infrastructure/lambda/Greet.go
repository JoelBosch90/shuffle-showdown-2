package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

type IGreet interface {
	Create(stack awscdk.Stack) awslambda.Function
}

type Greet struct{}

func (g *Greet) Create(stack awscdk.Stack) awslambda.Function {
	lambda := awslambda.NewFunction(stack, jsii.String("GreetFunction"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("../../controllers/greet"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	gateway := awsapigateway.NewLambdaRestApi(stack, jsii.String("HelloWorldGateway"), &awsapigateway.LambdaRestApiProps{
		Handler: lambda,
		Proxy:   jsii.Bool(true),
	})

	helloResource := gateway.Root().AddResource(jsii.String("hello"), &awsapigateway.ResourceOptions{})
	helloResource.AddMethod(jsii.String("GET"), awsapigateway.NewLambdaIntegration(lambda, &awsapigateway.LambdaIntegrationOptions{}), &awsapigateway.MethodOptions{})

	return lambda
}

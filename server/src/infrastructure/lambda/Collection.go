package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

type ICollection interface {
	Create(stack awscdk.Stack)
}

type Collection struct{}

func (c *Collection) Create(stack awscdk.Stack) {
	greet := &Greet{}
	greet.Create(stack)
}

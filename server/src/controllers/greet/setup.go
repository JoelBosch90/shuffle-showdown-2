package main

import (
	"greet/interfaces"
)

func setup(start interfaces.LambdaStarter) {
	start(handler)
}

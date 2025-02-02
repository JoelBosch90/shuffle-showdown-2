package main

import (
	"greet/interfaces"
)

func setup(start interfaces.StartLambda) {
	start(handler)
}

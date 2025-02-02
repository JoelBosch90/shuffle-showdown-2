package main

import (
	"greet/mocks"
	"testing"
)

func TestSetup(t *testing.T) {

	t.Run("greet.setup", func(t *testing.T) {
		// SETUP
		t.Parallel()

		mockLambdaStarter := mocks.StartLambda{}
		mockLambdaStarter.SetT(t)

		// WHEN
		setup(mockLambdaStarter.Get())

		// THEN
		if mockLambdaStarter.TimesCalled() < 1 {
			t.Errorf("Start callback was never called.")
		}

		if mockLambdaStarter.TimesCalled() > 1 {
			t.Errorf("Start callback was called too many times.")
		}
	})
}

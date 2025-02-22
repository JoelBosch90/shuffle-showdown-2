package main

import (
	"greet/interfaces"
	"greet/mocks"
	"reflect"
	"testing"
)

type MockStartLambda struct {
	mocks.Function
}

func (m *MockStartLambda) Get() interfaces.StartLambda {
	return func(handler interface{}) {
		m.SetTimesCalled(m.TimesCalled() + 1)
		callbackType := reflect.TypeOf(handler)

		if callbackType.Kind() != reflect.Func {
			m.T().Errorf("Expected callback to be a function, got %v", callbackType.Kind())
			return
		}

		if callbackType.NumOut() != 2 {
			m.T().Errorf("Expected callback to have 2 return values, got %d", callbackType.NumOut())
			return
		}
	}
}

func TestSetup(t *testing.T) {

	t.Run("greet.setup", func(t *testing.T) {
		// SETUP
		t.Parallel()

		mockLambdaStarter := MockStartLambda{}
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

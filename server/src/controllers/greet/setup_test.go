package main

import (
	"reflect"
	"testing"
)

func TestSetup(t *testing.T) {
	t.Run("greet.setup", func(t *testing.T) {
		// SETUP
		t.Parallel()

		// GIVEN
		isCalled := false
		given := func(callback interface{}) {
			isCalled = true
			callbackType := reflect.TypeOf(callback)

			if callbackType.Kind() != reflect.Func {
				t.Errorf("Expected callback to be a function, got %v", callbackType.Kind())
				return
			}

			if callbackType.NumOut() != 2 {
				t.Errorf("Expected callback to have 2 return values, got %d", callbackType.NumOut())
				return
			}
		}

		// WHEN
		setup(given)

		// THEN
		if !isCalled {
			t.Errorf("Start callback was never called.")
		}
	})
}

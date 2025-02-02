package mocks

import (
	"greet/interfaces"
	"reflect"
)

type LambdaStarter struct {
	Function
}

func (m *LambdaStarter) GetFunction() interfaces.LambdaStarter {
	return func(handler interface{}) {
		m.timesCalled += 1
		callbackType := reflect.TypeOf(handler)

		if callbackType.Kind() != reflect.Func {
			m.t.Errorf("Expected callback to be a function, got %v", callbackType.Kind())
			return
		}

		if callbackType.NumOut() != 2 {
			m.t.Errorf("Expected callback to have 2 return values, got %d", callbackType.NumOut())
			return
		}
	}
}

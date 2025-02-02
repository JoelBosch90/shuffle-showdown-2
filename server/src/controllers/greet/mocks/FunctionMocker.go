package mocks

import "testing"

type FunctionMocker interface {
	SetT(t *testing.T)
	GetTimesCalled() int
	GetFunction() interface{}
}

type Function struct {
	t           *testing.T
	timesCalled int
}

func (m *Function) SetT(t *testing.T) {
	m.t = t
}

func (m *Function) GetTimesCalled() int {
	return m.timesCalled
}

func (m *Function) GetFunction() interface{} {
	return func(handler interface{}) interface{} {
		m.timesCalled++

		return nil
	}
}

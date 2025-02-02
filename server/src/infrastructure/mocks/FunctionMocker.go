package mocks

import "testing"

type FunctionMocker interface {
	SetT(t *testing.T)
	GetTimesCalled() int
	GetFunction() func()
}

type Function struct {
	timesCalled int
}

func (m *Function) GetTimesCalled() int {
	return m.timesCalled
}

func (m *Function) GetFunction() func() {
	return func() {
		m.timesCalled++
	}
}

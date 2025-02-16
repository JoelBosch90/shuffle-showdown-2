package mocks

import "testing"

type FunctionMocker interface {
	SetT(t *testing.T)
	GetTimesCalled() int
	GetFunction() func()
}

type Function struct {
	t           *testing.T
	timesCalled int
}

func (m *Function) T() *testing.T {
	return m.t
}

func (m *Function) SetT(t *testing.T) {
	m.t = t
}

func (m *Function) TimesCalled() int {
	return m.timesCalled
}

func (m *Function) SetTimesCalled(timesCalled int) {
	m.timesCalled = timesCalled
}

func (m *Function) Get() func() {
	return func() {
		m.timesCalled++
	}
}

package mocks

import context "context"

type MockContextKey string
type MockContextValue string

const (
	mockContextKey   MockContextKey   = "testKey"
	mockContextValue MockContextValue = "testValue"
)

func GetMockContext() context.Context {
	return context.WithValue(context.Background(), mockContextKey, mockContextValue)
}

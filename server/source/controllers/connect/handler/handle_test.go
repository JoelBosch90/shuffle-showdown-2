package handler

import (
	"testing"

	"connect/mocks"
)

func TestHandle(t *testing.T) {
	testCases := []struct {
		name                 string
		expectedStatusCode   int
		expectedBodyContains string
		expectError          bool
	}{
		{
			name:                 "returns the default error response",
			expectedStatusCode:   500,
			expectedBodyContains: "Failed to do anything",
			expectError:          true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockContext := mocks.GetMockContext()

			response, err := HandleWithDependencies(mockContext)

			if testCase.expectError && err == nil {
				t.Error("Expected error but got none")
				return
			}
			if !testCase.expectError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
				return
			}

			if response.StatusCode != testCase.expectedStatusCode {
				t.Errorf("Expected status code %d, got %d", testCase.expectedStatusCode, response.StatusCode)
			}

			if testCase.expectedBodyContains != "" && response.Body != testCase.expectedBodyContains {
				t.Errorf("Expected body to contain '%s', got '%s'", testCase.expectedBodyContains, response.Body)
			}
		})
	}
}

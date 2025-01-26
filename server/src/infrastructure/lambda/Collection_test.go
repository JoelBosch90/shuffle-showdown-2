package lambda

import (
	"testing"

	"infrastructure/mocks"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"go.uber.org/mock/gomock"
)

func TestCollection_Create(t *testing.T) {
	t.Run("lambda.Collection.Create", func(t *testing.T) {
		// SETUP
		t.Parallel()
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		// GIVEN
		mockStack := awscdk.NewStack(nil, nil, nil)
		mockGreet := mocks.NewMockIGreet(ctrl)

		collection := &Collection{}

		// WHEN
		collection.Create(mockStack)

		// THEN
		// Verify that Greet was called exactly once
		mockGreet.EXPECT().Create(mockStack).Times(1)
	})
}

package main

import (
	"infrastructure/interfaces"
	"infrastructure/mocks"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/golang/mock/gomock"
)

type mockNewTable struct {
	mocks.Function
}

func (m *mockNewTable) Get() interfaces.NewTable {
	return func(scope constructs.Construct, id *string, props *awsdynamodb.TableProps) awsdynamodb.Table {
		m.SetTimesCalled(m.TimesCalled() + 1)

		return nil
	}
}

func TestCreateTable(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	t.Run("interfaces.createTable", func(t *testing.T) {
		// SETUP
		t.Parallel()

		// GIVEN
		mockStack := mocks.NewMockStack(controller)
		mockTableParameters := interfaces.TableParameters{
			ID:               "HelloWorldTable",
			PartitionKeyName: "PK",
		}
		mockNewTable := mockNewTable{}

		// WHEN
		createTableWithDependencies(mockStack, mockTableParameters, mockNewTable.Get())

		// THEN
		newTableTimesCalled := mockNewTable.TimesCalled()
		if newTableTimesCalled != 1 {
			t.Errorf("Expected newTable to be called once, but was called %d times", newTableTimesCalled)
		}
	})
}

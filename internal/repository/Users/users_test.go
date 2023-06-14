package users

import (
	"context"
	"errors"
	connectionTest "pdi-go-kafka-bd/internal/database/mock"
	"pdi-go-kafka-bd/internal/entity"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsersRepository(t *testing.T) {

	pgConnection := connectionTest.NewPgConnection(t)
	createRepository := CreateRepository(pgConnection)
	assert.NotNil(t, createRepository)

}

func TestCreateUserReturnSuccess(t *testing.T) {
	ctx := context.Background()
	pgConnection := connectionTest.NewPgConnection(t)
	createRepository := CreateRepository(pgConnection)
	mockUser := entity.User{
		Name: "Test for Mock",
	}
	pgConnection.On("Query",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.Anything,
		mock.Anything).Return(
		&connectionTest.MockRows{
			FillInterfaceOnScan: func(dest *[]interface{}) error {
				Id := (*dest)[0].(*string)
				Name := (*dest)[1].(*string)

				*Id = uuid.New().String()
				*Name = mockUser.Name
				return nil
			},
			RowCounter: 1,
		},
		nil,
	)

	resp, err := createRepository.InsertUser(ctx, mockUser)
	assert.Nil(t, err)
	assert.NotNil(t, resp.Name)
}

func TestCreateUserReturnFail(t *testing.T) {
	ctx := context.Background()
	pgConnection := connectionTest.NewPgConnection(t)
	createRepository := CreateRepository(pgConnection)
	mockUser := entity.User{
		Name: "Test for Mock",
	}
	pgConnection.On("Query",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.Anything,
		mock.Anything).Return(
		&connectionTest.MockRows{
			FillInterfaceOnScan: func(dest *[]interface{}) error { return nil },
			RowCounter:          1,
		},
		errors.New("Query execution error"),
	)

	resp, err := createRepository.InsertUser(ctx, mockUser)
	assert.Error(t, err)
	assert.Nil(t, resp)

}

func TestCreateUserReturnEmpty(t *testing.T) {
	ctx := context.Background()
	pgConnection := connectionTest.NewPgConnection(t)
	createRepository := CreateRepository(pgConnection)
	mockUser := entity.User{
		Name: "Test for Mock",
	}
	pgConnection.On("Query",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.Anything,
		mock.Anything).Return(
		&connectionTest.MockRows{
			FillInterfaceOnScan: func(dest *[]interface{}) error { return errors.New("Error in Scan") },
			RowCounter:          1,
		},
		nil,
	)

	resp, err := createRepository.InsertUser(ctx, mockUser)
	assert.Error(t, err)
	assert.Empty(t, resp)

}

func TestGetUsersReturnSuccess(t *testing.T) {
	ctx := context.Background()
	pgConnection := connectionTest.NewPgConnection(t)
	createRepository := CreateRepository(pgConnection)

	pgConnection.On("Query",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.Anything,
		mock.Anything,
		mock.Anything).
		Return(
			&connectionTest.MockRows{
				FillInterfaceOnScan: func(dest *[]interface{}) error {
					Id := (*dest)[0].(*string)
					Name := (*dest)[1].(*string)

					*Id = uuid.New().String()
					*Name = "Test with mock"
					return nil
				},
				RowCounter: 1,
			},
			nil,
		)
	resp, err := createRepository.GetUsers(ctx)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
func TestGetUsersReturnError(t *testing.T) {
	ctx := context.Background()
	pgConnection := connectionTest.NewPgConnection(t)
	createRepository := CreateRepository(pgConnection)

	pgConnection.On("Query",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(
		&connectionTest.MockRows{
			FillInterfaceOnScan: func(dest *[]interface{}) error { return nil },
			RowCounter:          1,
		},
		errors.New("Query execution error"),
	)
	resp, err := createRepository.GetUsers(ctx)
	assert.Error(t, err)
	assert.Nil(t, resp)
}
func TestGetUsersReturnEmpty(t *testing.T) {
	ctx := context.Background()
	pgConnection := connectionTest.NewPgConnection(t)
	createRepository := CreateRepository(pgConnection)

	pgConnection.On("Query",
		mock.AnythingOfType("*context.emptyCtx"),
		mock.Anything,
		mock.Anything,
		mock.Anything).Return(
		&connectionTest.MockRows{
			FillInterfaceOnScan: func(dest *[]interface{}) error { return errors.New("Error in Scan") },
			RowCounter:          1,
		},
		nil,
	)

	resp, err := createRepository.GetUsers(ctx)
	assert.Error(t, err)
	assert.Empty(t, resp)
}

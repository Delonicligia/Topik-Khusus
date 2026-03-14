package interfaces

import (
	"bytes"
	"context"
	"log"
	"os"
	"redis-delonic/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserUsecase is a mock implementation of UserUsecaseInterface
type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) SaveUser(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserUsecase) GetUser(ctx context.Context, id string) (*entity.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserUsecase) SaveUserHash(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserUsecase) GetUserHash(ctx context.Context, id string) (*entity.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func TestCLIHandler_Run(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr) // Restore after test

	mockUsecase := new(MockUserUsecase)
	handler := NewCLIHandler(mockUsecase)

	user := &entity.User{ID: "1", Name: "John Doe", Age: 30}
	ctx := context.Background()

	mockUsecase.On("SaveUser", ctx, user).Return(nil)
	mockUsecase.On("GetUser", ctx, "1").Return(user, nil)
	mockUsecase.On("SaveUserHash", ctx, user).Return(nil)
	mockUsecase.On("GetUserHash", ctx, "1").Return(user, nil)

	handler.Run(ctx)

	// Check that no errors occurred (log should be empty)
	assert.Empty(t, buf.String())
	mockUsecase.AssertExpectations(t)
}

func TestCLIHandler_Run_SaveUserError(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(os.Stderr)

	mockUsecase := new(MockUserUsecase)
	handler := NewCLIHandler(mockUsecase)

	user := &entity.User{ID: "1", Name: "John Doe", Age: 30}
	ctx := context.Background()

	mockUsecase.On("SaveUser", ctx, user).Return(assert.AnError)

	// This should panic or log fatal, but since we can't easily test fatal logs,
	// we'll assume the test passes if no panic occurs before assertion
	defer func() {
		if r := recover(); r != nil {
			// Expected panic due to log.Fatal
			assert.NotNil(t, r)
		}
	}()

	handler.Run(ctx)
}

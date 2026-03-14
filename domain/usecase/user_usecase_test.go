package usecase

import (
	"context"
	"errors"
	"redis-delonic/domain/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) SaveUser(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUser(ctx context.Context, id string) (*entity.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *MockUserRepository) SaveUserHash(ctx context.Context, user *entity.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) GetUserHash(ctx context.Context, id string) (*entity.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
}

func TestUserUsecase_SaveUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	user := &entity.User{ID: "1", Name: "John", Age: 30}
	ctx := context.Background()

	mockRepo.On("SaveUser", ctx, user).Return(nil)

	err := usecase.SaveUser(ctx, user)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_GetUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	user := &entity.User{ID: "1", Name: "John", Age: 30}
	ctx := context.Background()

	mockRepo.On("GetUser", ctx, "1").Return(user, nil)

	result, err := usecase.GetUser(ctx, "1")

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_SaveUserHash(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	user := &entity.User{ID: "1", Name: "John", Age: 30}
	ctx := context.Background()

	mockRepo.On("SaveUserHash", ctx, user).Return(nil)

	err := usecase.SaveUserHash(ctx, user)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_GetUserHash(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	user := &entity.User{ID: "1", Name: "John", Age: 30}
	ctx := context.Background()

	mockRepo.On("GetUserHash", ctx, "1").Return(user, nil)

	result, err := usecase.GetUserHash(ctx, "1")

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_SaveUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	user := &entity.User{ID: "1", Name: "John", Age: 30}
	ctx := context.Background()

	mockRepo.On("SaveUser", ctx, user).Return(errors.New("save error"))

	err := usecase.SaveUser(ctx, user)

	assert.Error(t, err)
	assert.Equal(t, "save error", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_GetUser_Error(t *testing.T) {
	mockRepo := new(MockUserRepository)
	usecase := NewUserUsecase(mockRepo)

	ctx := context.Background()

	mockRepo.On("GetUser", ctx, "1").Return(nil, errors.New("get error"))

	result, err := usecase.GetUser(ctx, "1")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "get error", err.Error())
	mockRepo.AssertExpectations(t)
}

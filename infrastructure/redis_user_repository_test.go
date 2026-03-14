package infrastructure

import (
	"context"
	"encoding/json"
	"errors"
	"redis-delonic/domain/entity"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRedisClient is a mock implementation of RedisClientInterface
type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	args := m.Called(ctx, key, value, expiration)
	cmd := redis.NewStatusCmd(ctx)
	if args.Error(0) != nil {
		cmd.SetErr(args.Error(0))
	} else {
		cmd.SetVal("OK")
	}
	return cmd
}

func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	args := m.Called(ctx, key)
	cmd := redis.NewStringCmd(ctx)
	if args.Error(1) != nil {
		cmd.SetErr(args.Error(1))
	} else {
		cmd.SetVal(args.String(0))
	}
	return cmd
}

func (m *MockRedisClient) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	args := m.Called(ctx, key, values)
	cmd := redis.NewIntCmd(ctx)
	if args.Error(0) != nil {
		cmd.SetErr(args.Error(0))
	} else {
		cmd.SetVal(1)
	}
	return cmd
}

func (m *MockRedisClient) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	args := m.Called(ctx, key, field)
	cmd := redis.NewStringCmd(ctx)
	if args.Error(1) != nil {
		cmd.SetErr(args.Error(1))
	} else {
		cmd.SetVal(args.String(0))
	}
	return cmd
}

func TestRedisUserRepository_SaveUser(t *testing.T) {
	mockClient := new(MockRedisClient)
	repo := NewRedisUserRepository(mockClient).(*RedisUserRepository)

	user := &entity.User{ID: "1", Name: "John", Age: 30}
	ctx := context.Background()

	expectedKey := "user:1"

	mockClient.On("Set", ctx, expectedKey, mock.AnythingOfType("[]uint8"), time.Duration(0)).Return(nil)

	err := repo.SaveUser(ctx, user)

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestRedisUserRepository_GetUser(t *testing.T) {
	mockClient := new(MockRedisClient)
	repo := NewRedisUserRepository(mockClient).(*RedisUserRepository)

	user := &entity.User{ID: "1", Name: "John", Age: 30}
	ctx := context.Background()

	expectedKey := "user:1"
	userJSON, _ := json.Marshal(user)

	mockClient.On("Get", ctx, expectedKey).Return(string(userJSON), nil)

	result, err := repo.GetUser(ctx, "1")

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockClient.AssertExpectations(t)
}

func TestRedisUserRepository_SaveUserHash(t *testing.T) {
	mockClient := new(MockRedisClient)
	repo := NewRedisUserRepository(mockClient).(*RedisUserRepository)

	user := &entity.User{ID: "1", Name: "John", Age: 30}
	ctx := context.Background()

	expectedKey := "user_hash:1"

	mockClient.On("HSet", ctx, expectedKey, "name", "John", "age", 30).Return(nil)

	err := repo.SaveUserHash(ctx, user)

	assert.NoError(t, err)
	mockClient.AssertExpectations(t)
}

func TestRedisUserRepository_GetUserHash(t *testing.T) {
	mockClient := new(MockRedisClient)
	repo := NewRedisUserRepository(mockClient).(*RedisUserRepository)

	ctx := context.Background()

	expectedKey := "user_hash:1"

	mockClient.On("HGet", ctx, expectedKey, "name").Return("John", nil)
	mockClient.On("HGet", ctx, expectedKey, "age").Return("30", nil)

	result, err := repo.GetUserHash(ctx, "1")

	assert.NoError(t, err)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "John", result.Name)
	assert.Equal(t, 30, result.Age)
	mockClient.AssertExpectations(t)
}

func TestRedisUserRepository_SaveUser_Error(t *testing.T) {
	mockClient := new(MockRedisClient)
	repo := NewRedisUserRepository(mockClient).(*RedisUserRepository)

	user := &entity.User{ID: "1", Name: "John", Age: 30}
	ctx := context.Background()

	expectedKey := "user:1"

	mockClient.On("Set", ctx, expectedKey, mock.AnythingOfType("[]uint8"), time.Duration(0)).Return(errors.New("redis error"))

	err := repo.SaveUser(ctx, user)

	assert.Error(t, err)
	assert.Equal(t, "redis error", err.Error())
	mockClient.AssertExpectations(t)
}

func TestRedisUserRepository_GetUser_Error(t *testing.T) {
	mockClient := new(MockRedisClient)
	repo := NewRedisUserRepository(mockClient).(*RedisUserRepository)

	ctx := context.Background()

	expectedKey := "user:1"

	mockClient.On("Get", ctx, expectedKey).Return("", errors.New("redis error"))

	result, err := repo.GetUser(ctx, "1")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, "redis error", err.Error())
	mockClient.AssertExpectations(t)
}

func TestRedisUserRepository_GetUserHash_InvalidAge(t *testing.T) {
	mockClient := new(MockRedisClient)
	repo := NewRedisUserRepository(mockClient).(*RedisUserRepository)

	ctx := context.Background()

	expectedKey := "user_hash:1"

	mockClient.On("HGet", ctx, expectedKey, "name").Return("John", nil)
	mockClient.On("HGet", ctx, expectedKey, "age").Return("invalid", nil)

	result, err := repo.GetUserHash(ctx, "1")

	assert.Error(t, err)
	assert.Nil(t, result)
	mockClient.AssertExpectations(t)
}

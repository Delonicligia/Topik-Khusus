package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"redis-delonic/domain/entity"
	"redis-delonic/domain/repository"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClientInterface defines the interface for Redis operations
type RedisClientInterface interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	HGet(ctx context.Context, key, field string) *redis.StringCmd
}

// RedisUserRepository implements UserRepository using Redis
type RedisUserRepository struct {
	client RedisClientInterface
}

// NewRedisUserRepository creates a new RedisUserRepository
func NewRedisUserRepository(client RedisClientInterface) repository.UserRepository {
	return &RedisUserRepository{
		client: client,
	}
}

// SaveUser saves a user as JSON string
func (r *RedisUserRepository) SaveUser(ctx context.Context, user *entity.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, fmt.Sprintf("user:%s", user.ID), data, 0).Err()
}

// GetUser retrieves a user from JSON string
func (r *RedisUserRepository) GetUser(ctx context.Context, id string) (*entity.User, error) {
	data, err := r.client.Get(ctx, fmt.Sprintf("user:%s", id)).Result()
	if err != nil {
		return nil, err
	}
	var user entity.User
	err = json.Unmarshal([]byte(data), &user)
	return &user, err
}

// SaveUserHash saves a user as hash
func (r *RedisUserRepository) SaveUserHash(ctx context.Context, user *entity.User) error {
	key := fmt.Sprintf("user_hash:%s", user.ID)
	return r.client.HSet(ctx, key, "name", user.Name, "age", user.Age).Err()
}

// GetUserHash retrieves a user from hash
func (r *RedisUserRepository) GetUserHash(ctx context.Context, id string) (*entity.User, error) {
	key := fmt.Sprintf("user_hash:%s", id)
	name, err := r.client.HGet(ctx, key, "name").Result()
	if err != nil {
		return nil, err
	}
	ageStr, err := r.client.HGet(ctx, key, "age").Result()
	if err != nil {
		return nil, err
	}
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:   id,
		Name: name,
		Age:  age,
	}, nil
}

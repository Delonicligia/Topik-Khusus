package repository

import (
	"context"
	"redis-delonic/domain/entity"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	SaveUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, id string) (*entity.User, error)
	SaveUserHash(ctx context.Context, user *entity.User) error
	GetUserHash(ctx context.Context, id string) (*entity.User, error)
}

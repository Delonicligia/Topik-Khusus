package usecase

import (
	"context"
	"redis-delonic/domain/entity"
	"redis-delonic/domain/repository"
)

// UserUsecaseInterface defines the interface for user usecase
type UserUsecaseInterface interface {
	SaveUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, id string) (*entity.User, error)
	SaveUserHash(ctx context.Context, user *entity.User) error
	GetUserHash(ctx context.Context, id string) (*entity.User, error)
}

// UserUsecase defines the business logic for user operations
type UserUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase creates a new UserUsecase
func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

// SaveUser saves a user using string key
func (u *UserUsecase) SaveUser(ctx context.Context, user *entity.User) error {
	return u.userRepo.SaveUser(ctx, user)
}

// GetUser retrieves a user using string key
func (u *UserUsecase) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return u.userRepo.GetUser(ctx, id)
}

// SaveUserHash saves a user using hash
func (u *UserUsecase) SaveUserHash(ctx context.Context, user *entity.User) error {
	return u.userRepo.SaveUserHash(ctx, user)
}

// GetUserHash retrieves a user using hash
func (u *UserUsecase) GetUserHash(ctx context.Context, id string) (*entity.User, error) {
	return u.userRepo.GetUserHash(ctx, id)
}

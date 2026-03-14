package interfaces

import (
	"context"
	"fmt"
	"log"
	"redis-delonic/domain/entity"
	"redis-delonic/domain/usecase"
)

// CLIHandler handles CLI interactions
type CLIHandler struct {
	userUsecase usecase.UserUsecaseInterface
}

// NewCLIHandler creates a new CLIHandler
func NewCLIHandler(userUsecase usecase.UserUsecaseInterface) *CLIHandler {
	return &CLIHandler{
		userUsecase: userUsecase,
	}
}

// Run executes the CLI demo
func (h *CLIHandler) Run(ctx context.Context) {
	// Create a sample user
	user := &entity.User{
		ID:   "1",
		Name: "John Doe",
		Age:  30,
	}

	// Save user as string
	err := h.userUsecase.SaveUser(ctx, user)
	if err != nil {
		log.Fatal("Failed to save user:", err)
	}
	fmt.Println("User saved as string")

	// Get user
	retrievedUser, err := h.userUsecase.GetUser(ctx, "1")
	if err != nil {
		log.Fatal("Failed to get user:", err)
	}
	fmt.Printf("Retrieved user: ID=%s, Name=%s, Age=%d\n", retrievedUser.ID, retrievedUser.Name, retrievedUser.Age)

	// Save user as hash
	err = h.userUsecase.SaveUserHash(ctx, user)
	if err != nil {
		log.Fatal("Failed to save user hash:", err)
	}
	fmt.Println("User saved as hash")

	// Get user from hash
	retrievedUserHash, err := h.userUsecase.GetUserHash(ctx, "1")
	if err != nil {
		log.Fatal("Failed to get user hash:", err)
	}
	fmt.Printf("Retrieved user from hash: ID=%s, Name=%s, Age=%d\n", retrievedUserHash.ID, retrievedUserHash.Name, retrievedUserHash.Age)
}

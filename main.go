package main

import (
	"context"
	"redis-delonic/domain/usecase"
	"redis-delonic/infrastructure"
	"redis-delonic/interfaces"

	"github.com/redis/go-redis/v9"
)

func main() {
	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Initialize repository
	userRepo := infrastructure.NewRedisUserRepository(rdb)

	// Initialize usecase
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Initialize handler
	cliHandler := interfaces.NewCLIHandler(userUsecase)

	// Run the demo
	ctx := context.Background()
	cliHandler.Run(ctx)

	// Close connection
	rdb.Close()
}

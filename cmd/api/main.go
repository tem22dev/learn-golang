package main

import (
	"learn-golang/internal/config"
	"learn-golang/internal/handler"
	"learn-golang/internal/repository"
	"learn-golang/internal/routes"
	"learn-golang/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init config
	cfg := config.NewConfig()

	// Init repository
	userRepo := repository.NewInMemoryUserRepository()

	// Init service
	userService := service.NewUserService(userRepo)

	// Init handler
	userHandler := handler.NewUserHandler(userService)

	// Init routes
	userRoutes := routes.NewUserRoutes(userHandler)

	r := gin.Default()

	routes.RegisterRoutes(r, userRoutes)

	if err := r.Run(cfg.ServerAddress); err != nil {
		panic(err)
	}
}

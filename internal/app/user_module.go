package app

import (
	"learn-golang/internal/handler"
	"learn-golang/internal/repository"
	"learn-golang/internal/routes"
	"learn-golang/internal/service"
)

type UserModule struct {
	routes routes.Route
}

func NewUserModule() *UserModule {
	// Init repository
	userRepo := repository.NewInMemoryUserRepository()

	// Init service
	userService := service.NewUserService(userRepo)

	// Init handler
	userHandler := handler.NewUserHandler(userService)

	// Init routes
	userRoutes := routes.NewUserRoutes(userHandler)

	return &UserModule{
		routes: userRoutes,
	}
}

func (m *UserModule) Routes() routes.Route {
	return m.routes
}

package app

import (
	"learn-golang/internal/config"
	"learn-golang/internal/routes"

	"github.com/gin-gonic/gin"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config *config.Config
	router *gin.Engine
}

func NewApplication(cfg *config.Config) *Application {
	r := gin.Default()

	modules := []Module{
		NewUserModule(),
	}

	routes.RegisterRoutes(r, getModuleRoutes(modules)...)

	return &Application{
		config: cfg,
		router: r,
	}
}

func (a *Application) Run() error {
	return a.router.Run(a.config.ServerAddress)
}

func getModuleRoutes(modules []Module) []routes.Route {
	routeList := make([]routes.Route, len(modules))
	for i, module := range modules {
		routeList[i] = module.Routes()
	}

	return routeList
}

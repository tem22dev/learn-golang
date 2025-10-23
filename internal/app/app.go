package app

import (
	"learn-golang/internal/config"
	"learn-golang/internal/routes"
	"learn-golang/internal/validation"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

type Module interface {
	Routes() routes.Route
}

type Application struct {
	config  *config.Config
	router  *gin.Engine
	modules []Module
}

func NewApplication(cfg *config.Config) *Application {
	r := gin.Default()

	if err := validation.InitValidator(); err != nil {
		return nil
	}

	loadEnv()

	modules := []Module{
		NewUserModule(),
	}

	routes.RegisterRoutes(r, getModuleRoutes(modules)...)

	return &Application{
		config:  cfg,
		router:  r,
		modules: modules,
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

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found")
	}
}

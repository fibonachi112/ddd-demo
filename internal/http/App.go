package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"repo/config"
	"repo/internal/DI"
)

type HttpApp struct {
	FiberApp *fiber.App
	Config   *config.Configuration
	Di       *DI.ServiceContainer
}

func Make(configuration *config.Configuration, di *DI.ServiceContainer) HttpApp {
	app := HttpApp{}
	app.Config = configuration
	app.Di = di

	return app
}

func (a *HttpApp) initApp() {
	a.FiberApp = fiber.New(fiber.Config{
		AppName: a.Config.HttpApp.AppName,
	})
}

func (a *HttpApp) Serve() {
	a.initApp()
	a.makeRoutes()
	err := a.FiberApp.Listen(fmt.Sprintf("%s%s", a.Config.HttpApp.ListenAddr, a.Config.HttpApp.ListenPort))
	if err != nil {
		fmt.Print("app listen error occurred")
	}
}

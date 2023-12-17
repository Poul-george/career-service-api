package router

import (
	"github.com/Poul-george/go-api/api/config"
	"github.com/Poul-george/go-api/api/presentation/appapi/echoserver/inject"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouterAuthApi(g *echo.Group) {
	injector := inject.NewInjector(
		config.GetMySQLConfig(),
		config.GetServerConfig(),
		config.GetRedisConfig(),
		config.GetServiceConfig(),
	)
	routerAuthApi(g, injector)
}

func routerAuthApi(g *echo.Group, injector inject.Injector) {
	g.Use(
		middleware.CORS(),
	)
	g.POST("/login", injector.AuthLoginController())
}

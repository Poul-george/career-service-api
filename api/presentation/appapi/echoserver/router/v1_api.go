package router

import (
	"github.com/Poul-george/go-api/api/config"
	"github.com/Poul-george/go-api/api/presentation/appapi/echoserver/inject"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouterV1Api(g *echo.Group) {
	injector := inject.NewInjector(
		config.GetMySQLConfig(),
		config.GetServerConfig(),
	)
	routerV1Api(g, injector)
}

func routerV1Api(g *echo.Group, injector inject.Injector) {
	g.Use(
		middleware.CORS(),
	)
	g.GET("/users", injector.V1UserListController())
	g.GET("/users/detail", injector.V1UserDetailController())
	g.POST("/users", injector.V1UserCreateController())
}

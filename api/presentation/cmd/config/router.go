package config

import (
	appapi "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/router"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func router(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	appapi.RouterV1Api(e.Group("api/v1"))
}

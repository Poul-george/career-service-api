package config

//echoのcros設定のファイル airでの実装を考慮

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func NewEchoServer() *echo.Echo {
	e := SetMiddlleware()
	router(e)

	return e
}

func SetMiddlleware() *echo.Echo {
	e := echo.New()

	//CORS用のmiddlewareはあるものの、403を勝手に返してくれるわけではない。
	e.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			AllowHeaders: []string{
				echo.HeaderAccessControlAllowHeaders,
				echo.HeaderContentType,
				echo.HeaderContentLength,
				echo.HeaderAcceptEncoding,
				echo.HeaderXCSRFToken,
				echo.HeaderAuthorization,
			},
			AllowMethods: []string{
				http.MethodGet,
				http.MethodPut,
				http.MethodPatch,
				http.MethodPost,
				http.MethodDelete,
			},
			MaxAge: 86400,
		}),
	)

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	return e
}

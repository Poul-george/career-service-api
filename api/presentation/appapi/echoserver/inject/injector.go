package inject

import (
	"github.com/Poul-george/go-api/api/config"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/handler"
	echoContext "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"
	"github.com/labstack/echo/v4"
)

type Injector struct {
	mySQLConfig  config.MySQL
	serverConfig config.Server
}

func NewInjector(
	mySQLConfig config.MySQL,
	serverConfig config.Server,
) Injector {
	return Injector{
		mySQLConfig:  mySQLConfig,
		serverConfig: serverConfig,
	}
}

type callFunc func(ctx echoContext.Context) error

func newHandlerFunc(h callFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// return h(c.(*echoContext.Context))
		// 上のやり方だとなぜかうまくいかないので、後で原因をしっかり調べる
		ctx := echoContext.Context{Context: c}
		return h(ctx)
	}
}

func (i *Injector) gormHandler() *handler.Handler {
	return handler.NewHandler(i.mySQLConfig)
}

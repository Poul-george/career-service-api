package inject

import (
	"github.com/Poul-george/go-api/api/config"
	cacheHandler "github.com/Poul-george/go-api/api/infrastructure/data/persistence/cache/handler"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/handler"
	echoContext "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"
	"github.com/labstack/echo/v4"
)

type Injector struct {
	mySQLConfig   config.MySQL
	serverConfig  config.Server
	redisConfig   config.Redis
	serviceConfig config.Service
}

func NewInjector(
	mySQLConfig config.MySQL,
	serverConfig config.Server,
	redisConfig config.Redis,
	serviceConfig config.Service,
) Injector {
	return Injector{
		mySQLConfig:   mySQLConfig,
		serverConfig:  serverConfig,
		redisConfig:   redisConfig,
		serviceConfig: serviceConfig,
	}
}

type callFunc func(ctx echoContext.Context) error

func newHandlerFunc(h callFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := echoContext.Context{Context: c}
		return h(ctx)
	}
}

func (i *Injector) gormHandler() *handler.Handler {
	return handler.NewHandler(i.mySQLConfig)
}

func (i *Injector) cacheHandler() *cacheHandler.RedisHandler {
	return cacheHandler.NewCacheHandler(i.redisConfig)
}

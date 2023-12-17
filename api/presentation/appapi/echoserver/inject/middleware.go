package inject

import (
	"github.com/Poul-george/go-api/api/config"
	"github.com/Poul-george/go-api/api/core/domain/repository/auth/token"
	"github.com/labstack/echo/v4"
)

type verifyJWTFunc func(service config.Service, repository token.Repository) echo.MiddlewareFunc

func (i *Injector) VerifyJWTMiddleware(f verifyJWTFunc) echo.MiddlewareFunc {
	return f(i.serviceConfig, i.tokenRepository())
}

package router

import (
	"github.com/Poul-george/go-api/api/config"
	"github.com/Poul-george/go-api/api/core/domain/repository/auth/token"
	"github.com/Poul-george/go-api/api/presentation/appapi/common/response"
	"github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"
	"github.com/Poul-george/go-api/api/util/errors"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func verifyJWT(c config.Service, r token.Repository) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		ParseTokenFunc: func(ctx echo.Context, authToken string) (any, error) {
			return r.FindUserIDByAccessToken(ctx.Request().Context(), authToken, c.SecretKey)
		},
		ContextKey: context.UserIDKey(),
		ErrorHandler: func(ctx echo.Context, err error) error {
			return response.Unauthorized(context.Context{Context: ctx}, errors.New("authorization required"))
		},
	})
}

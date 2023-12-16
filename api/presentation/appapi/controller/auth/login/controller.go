package login

import (
	"github.com/Poul-george/go-api/api/core/common/types/identifier"
	usecase "github.com/Poul-george/go-api/api/core/usecase/api/auth/login"
	"github.com/Poul-george/go-api/api/presentation/appapi/common/response"
	echoContext "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"
	"github.com/Poul-george/go-api/api/util/errors"
	"net/http"
)

type Parameter struct {
	UserID         uint64 `json:"user_id"`
	ExternalUserID string `json:"external_user_id"`
}

type Controller struct {
	UseCase usecase.UseCase
}

func NewController(
	useCase usecase.UseCase,
) Controller {
	return Controller{
		UseCase: useCase,
	}
}

func (c Controller) Post(ctx echoContext.Context) error {
	var p Parameter
	if err := ctx.Bind(&p); err != nil {
		return response.BadRequest(ctx, err)
	}

	if p.ExternalUserID == "" && p.UserID == 0 {
		return response.BadRequest(ctx, errors.New("リクエストが不正です"))
	}

	out, err := c.UseCase.Do(ctx.Request().Context(), usecase.Input{
		ExternalUserID: identifier.ExternalUserID(p.ExternalUserID),
		UserID:         identifier.UserID(p.UserID),
	})
	if err != nil {
		return response.BadRequest(ctx, errors.New(err.Error()))
	}

	return response.OKWithItem(ctx, http.StatusOK, newResponse(out))
}

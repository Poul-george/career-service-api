package list

import (
	"github.com/Poul-george/go-api/api/presentation/appapi/common/response"
	echoContext "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"
	"github.com/Poul-george/go-api/api/util/errors"

	usecase "github.com/Poul-george/go-api/api/core/usecase/api/user/list"
)

type Parameter struct {
	UserID         *uint64 `json:"user_id" form:"user_id"`
	ExternalUserID *string `json:"external_user_id" form:"external_user_id"`
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

func (c Controller) Get(ctx echoContext.Context) error {
	var p Parameter
	if err := ctx.Bind(&p); err != nil {
		return err
	}

	out, err := c.UseCase.Do(ctx.Request().Context())
	if err != nil {
		return response.BadRequest(ctx, errors.New("リクエストができませんでした"))
	}

	res := NewResponse(*out)

	return response.OK(ctx, res)
}

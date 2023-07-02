package detail

import (
	"fmt"
	"net/http"

	"github.com/Poul-george/go-api/api/presentation/appapi/common/response"
	"github.com/Poul-george/go-api/api/util/errors"

	"github.com/Poul-george/go-api/api/core/common/types/identifier"

	echoContext "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"

	usecase "github.com/Poul-george/go-api/api/core/usecase/api/user/detail"
)

type Parameter struct {
	UserID         uint64 `json:"user_id" query:"user_id"`
	ExternalUserID string `json:"external_user_id" query:"external_user_id"`
}

type Controller struct {
	UseCase usecase.UseCase
}

func (c Controller) Get(ctx echoContext.Context) error {
	var p Parameter
	if err := ctx.Bind(&p); err != nil {
		fmt.Println("バインドエラー")
		return response.BadRequest(ctx, err)
	}

	fmt.Printf("============================ %s ============================= ", "detail get request")

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

	res := NewResponse(*out)

	return response.OKWithItem(ctx, http.StatusOK, res)
}

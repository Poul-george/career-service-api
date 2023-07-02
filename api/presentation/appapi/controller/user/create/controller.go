package create

import (
	"fmt"
	"github.com/Poul-george/go-api/api/presentation/appapi/common/response"
	"net/http"

	echoContext "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"

	"github.com/Poul-george/go-api/api/core/usecase/api/user/create"
	usecase "github.com/Poul-george/go-api/api/core/usecase/api/user/create"
)

type Parameter struct {
	ExternalUserID string `json:"external_user_id"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	MailAddress    string `json:"email"`
	Comments       string `json:"comments"`
}

type Controller struct {
	UseCase usecase.UseCase
}

func (c Controller) Post(ctx echoContext.Context) error {
	var p Parameter
	if err := ctx.Bind(&p); err != nil {
		return response.BadRequest(ctx, err)
	}

	fmt.Printf("============================ %v ============================= create parameter", p)

	err := c.UseCase.Do(ctx.Request().Context(), create.Input{
		ExternalUserID: p.ExternalUserID,
		Name:           p.Name,
		Password:       p.Password,
		MailAddress:    p.MailAddress,
		Comments:       p.Comments,
	})
	if err != nil {
		return response.InternalServerError(ctx, err)
	}

	return response.OKWithMessage(ctx, http.StatusOK, "正常に登録できました")
}

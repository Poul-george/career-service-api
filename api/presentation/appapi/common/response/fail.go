package response

import (
	echoContext "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"
	"net/http"
)

func BadRequest(c echoContext.Context, err error) error {
	return fail(c, http.StatusBadRequest, err.Error())
}

func Forbidden(c echoContext.Context, err error) error {
	return fail(c, http.StatusForbidden, err.Error())
}

func NotFound(c echoContext.Context, err error) error {
	return fail(c, http.StatusNotFound, err.Error())
}

// もう少し改良する
func InternalServerError(c echoContext.Context, err error) error {
	return fail(c, http.StatusInternalServerError, err.Error())
}

func fail(c echoContext.Context, code int, m string) error {
	return c.JSON(
		code,
		struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    -1,
			Message: m,
		},
	)
}

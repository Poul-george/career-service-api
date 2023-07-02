package response

import (
	echoContext "github.com/Poul-george/go-api/api/presentation/appapi/echoserver/context"
)

func OKWithItem(c echoContext.Context, code int, item interface{}) error {
	return c.JSON(
		code,
		struct {
			Code int         `json:"code"`
			Item interface{} `json:"item"`
		}{
			Code: 1,
			Item: item,
		},
	)
}

func OKWithMessage(c echoContext.Context, code int, m interface{}) error {
	return c.JSON(
		code,
		struct {
			Code    int         `json:"code"`
			Message interface{} `json:"message"`
		}{
			Code:    1,
			Message: m,
		},
	)
}

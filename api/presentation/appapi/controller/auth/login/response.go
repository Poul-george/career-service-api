package login

import "github.com/Poul-george/go-api/api/core/usecase/api/auth/login"

type Response struct {
	UserID         uint64 `json:"user_id"`
	ExternalUserID string `json:"external_user_id"`
	RefreshToken   string `json:"refresh_token"`
	AccessToken    string `json:"access_token"`
}

func newResponse(out *login.Output) *Response {
	return &Response{
		UserID:         out.UserID.Uint64(),
		ExternalUserID: out.ExternalUserID.String(),
		RefreshToken:   out.RefreshToken,
		AccessToken:    out.AccessToken,
	}
}

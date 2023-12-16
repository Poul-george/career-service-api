package login

import (
	"github.com/Poul-george/go-api/api/core/common/types/identifier"
	"github.com/Poul-george/go-api/api/core/domain/model"
)

type Output struct {
	UserID         identifier.UserID
	ExternalUserID identifier.ExternalUserID
	RefreshToken   string
	AccessToken    string
}

func newOutput(authToken model.AuthToken, accessToken string) Output {
	return Output{
		UserID:         authToken.UserID(),
		ExternalUserID: authToken.ExternalUserID(),
		RefreshToken:   authToken.RefreshToken().String(),
		AccessToken:    accessToken,
	}
}

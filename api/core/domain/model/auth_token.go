package model

import (
	"github.com/Poul-george/go-api/api/core/common/types/identifier"
	"github.com/Poul-george/go-api/api/core/domain/value"
	"github.com/Poul-george/go-api/api/util/localtime"
	"time"
)

type AuthToken struct {
	userID         identifier.UserID
	externalUserID identifier.ExternalUserID
	refreshToken   value.AuthTokenRefreshToken
	expiresAt      time.Time
}

func NewAuthToken(
	userID identifier.UserID,
	externalUserID identifier.ExternalUserID,
) *AuthToken {
	return &AuthToken{
		userID:         userID,
		externalUserID: externalUserID,
		refreshToken:   value.NewAuthTokenRefreshToken(userID, externalUserID),
		expiresAt:      localtime.Now().Add(30 * 24 * time.Hour), // 1ヶ月
	}
}

func (at *AuthToken) UserID() identifier.UserID {
	return at.userID
}
func (at *AuthToken) ExternalUserID() identifier.ExternalUserID {
	return at.externalUserID
}
func (at *AuthToken) RefreshToken() value.AuthTokenRefreshToken {
	return at.refreshToken
}
func (at *AuthToken) ExpiresAt() time.Time {
	return at.expiresAt
}

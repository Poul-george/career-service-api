package value

import (
	"fmt"
	"github.com/Poul-george/go-api/api/core/common/types/identifier"
	"github.com/google/uuid"
)

type AuthTokenRefreshToken string

func NewAuthTokenRefreshToken(userID identifier.UserID, externalUserID identifier.ExternalUserID) AuthTokenRefreshToken {
	return AuthTokenRefreshToken(fmt.Sprintf("%d:%s:%s", userID, externalUserID, uuid.New().String()))
}

func (s AuthTokenRefreshToken) String() string {
	return string(s)
}

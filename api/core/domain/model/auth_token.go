package model

import "github.com/Poul-george/go-api/api/core/common/types/identifier"

type AuthToken struct {
	userID         identifier.UserID
	externalUserID identifier.ExternalUserID
	grantType
}

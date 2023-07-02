package detail

import "github.com/Poul-george/go-api/api/core/common/types/identifier"

type Input struct {
	ExternalUserID identifier.ExternalUserID
	UserID         identifier.UserID
}

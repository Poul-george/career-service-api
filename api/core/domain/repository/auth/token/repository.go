package token

import (
	"context"
	"github.com/Poul-george/go-api/api/core/common/types/identifier"
	"github.com/Poul-george/go-api/api/core/domain/model"
)

type Repository interface {
	CreateAccessToken(context.Context, model.AuthToken, string) (*string, error)
	FindUserIDByAccessToken(context.Context, string, string) (*identifier.UserID, error)
}

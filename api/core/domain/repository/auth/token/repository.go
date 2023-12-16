package token

import (
	"context"
	"github.com/Poul-george/go-api/api/core/domain/model"
)

type Repository interface {
	CreateAccessToken(context.Context, model.AuthToken) (*string, error)
}

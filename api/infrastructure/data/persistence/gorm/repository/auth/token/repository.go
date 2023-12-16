package token

import (
	"context"
	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/core/domain/repository/auth/token"
	cacheHandler "github.com/Poul-george/go-api/api/infrastructure/data/persistence/cache/handler"
)

type Repository struct {
	handler *cacheHandler.RedisHandler
}

var _ token.Repository = (*Repository)(nil)

func NewRepository(h *cacheHandler.RedisHandler) *Repository {
	return &Repository{handler: h}
}

func (r *Repository) CreateAccessToken(ctx context.Context, authToken model.AuthToken, appKey string) (*string, error) {
	return generateJWT(authToken, appKey)
}

package user

import (
	"context"
	"github.com/Poul-george/go-api/api/core/common/types/identifier"

	"github.com/Poul-george/go-api/api/core/domain/model"
)

type Repository interface {
	Create(context.Context, *model.User) error
	FindByIDs(context.Context) ([]model.User, error)
	FindByID(context.Context, identifier.ExternalUserID, identifier.UserID) (*model.User, error)
}

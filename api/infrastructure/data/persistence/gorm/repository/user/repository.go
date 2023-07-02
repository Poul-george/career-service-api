package user

import (
	"context"
	"github.com/Poul-george/go-api/api/core/common/types/identifier"
	"github.com/Poul-george/go-api/api/util/errors"

	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/handler"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/table"
)

type Repository struct {
	handler *handler.Handler
}

func NewRepository(h *handler.Handler) *Repository {
	return &Repository{handler: h}
}

func (r *Repository) Create(ctx context.Context, user *model.User) error {
	createUser := ToUserForCreate(*user)
	if err := r.handler.Writer(ctx).Create(createUser).Error; err != nil {
		return errors.Wrap(err)
	}
	return nil
}

func (r *Repository) FindByIDs(ctx context.Context) ([]model.User, error) {
	users := []table.User{}
	err := r.handler.Reader(ctx).
		Model(&table.User{}).
		Order("id").
		Find(&users).Error
	if err != nil {
		return nil, errors.Wrap(err)
	}
	res := ToModelUsers(users)
	return res, nil
}

func (r *Repository) FindByID(
	ctx context.Context,
	externalUserID identifier.ExternalUserID,
	userID identifier.UserID,
) (*model.User, error) {
	user := table.User{}
	q := r.handler.Reader(ctx).Model(&table.User{})
	if externalUserID != "" {
		q = q.Where("external_user_id = ?", externalUserID)
	}
	if userID != 0 {
		q = q.Where("id = ?", userID)
	}
	err := q.First(&user).Error
	if err != nil {
		return nil, errors.Wrap(err)
	}

	res := ToModelUser(user)
	return res, nil
}

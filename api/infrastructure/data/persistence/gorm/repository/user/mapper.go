package user

import (
	"github.com/Poul-george/go-api/api/core/common/types/identifier"
	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/infrastructure/data/persistence/gorm/table"
)

func ToUserForCreate(u model.User) *table.User {
	t := table.User{
		ExternalUserID: u.ExternalUserID().String(),
		Name:           u.Name(),
		Password:       u.Password(),
		MailAddress:    u.MailAddress(),
		Comments:       u.Comments(),
	}
	return &t
}

func ToModelUser(t table.User) *model.User {
	return model.ReConstructorUser(
		identifier.UserID(t.ID),
		identifier.ExternalUserID(t.ExternalUserID),
		t.Name,
		t.Password,
		t.MailAddress,
		t.Comments,
		t.UpdatedAt,
	)
}

func ToModelUsers(t []table.User) []model.User {
	users := make([]model.User, len(t))
	for i, v := range t {
		u := ToModelUser(v)
		users[i] = *u
	}
	return users
}

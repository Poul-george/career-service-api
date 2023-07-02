package list

import (
	"github.com/Poul-george/go-api/api/core/domain/model"
	"time"
)

type Output struct {
	Users []OutputUser
}

type OutputUser struct {
	ID             uint64
	ExternalUserID string
	Name           string
	MailAddress    string
	Comments       string
	UpdatedAt      time.Time
}

func NewOutput(users []model.User) *Output {
	outputUsers := make([]OutputUser, len(users))

	for i, u := range users {
		outputUsers[i] = OutputUser{
			ID:             u.ID().Uint64(),
			ExternalUserID: u.ExternalUserID().String(),
			Name:           u.Name(),
			MailAddress:    u.MailAddress(),
			Comments:       u.Comments(),
			UpdatedAt:      u.UpdatedAt(),
		}
	}

	return &Output{Users: outputUsers}
}

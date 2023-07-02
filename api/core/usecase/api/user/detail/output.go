package detail

import (
	"github.com/Poul-george/go-api/api/core/domain/model"
	"time"
)

type Output struct {
	ID             uint64
	ExternalUserID string
	Name           string
	MailAddress    string
	Comments       string
	UpdatedAt      time.Time
}

func NewOutput(user model.User) *Output {
	return &Output{
		ID:             user.ID().Uint64(),
		ExternalUserID: user.ExternalUserID().String(),
		Name:           user.Name(),
		MailAddress:    user.MailAddress(),
		Comments:       user.Comments(),
		UpdatedAt:      user.UpdatedAt(),
	}
}

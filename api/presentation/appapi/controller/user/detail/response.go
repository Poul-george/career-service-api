package detail

import (
	"time"

	"github.com/Poul-george/go-api/api/core/usecase/api/user/detail"
)

type Response struct {
	ID             uint64    `json:"id"`
	ExternalUserID string    `json:"external_user_id"`
	Name           string    `json:"name"`
	MailAddress    string    `json:"mail_address"`
	Comments       string    `json:"comments"`
	UpdatedAt      time.Time `json:"latest_day"`
}

func NewResponse(user detail.Output) Response {
	return Response{
		ID:             user.ID,
		ExternalUserID: user.ExternalUserID,
		Name:           user.Name,
		MailAddress:    user.MailAddress,
		Comments:       user.Comments,
		UpdatedAt:      user.UpdatedAt,
	}
}

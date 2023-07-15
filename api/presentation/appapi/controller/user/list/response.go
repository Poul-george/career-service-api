package list

import (
	"time"

	"github.com/Poul-george/go-api/api/core/usecase/api/user/list"
)

type Response struct {
	Items []UserResponse `json:"items"`
}

type UserResponse struct {
	ID             uint64    `json:"id"`
	ExternalUserID string    `json:"externalUserId"`
	Name           string    `json:"name"`
	MailAddress    string    `json:"mailAddress"`
	Comments       string    `json:"comments"`
	UpdatedAt      time.Time `json:"latestDay"`
}

func NewResponse(out list.Output) Response {
	users := make([]UserResponse, len(out.Users))
	for i, u := range out.Users {
		users[i] = UserResponse{
			ID:             u.ID,
			ExternalUserID: u.ExternalUserID,
			Name:           u.Name,
			MailAddress:    u.MailAddress,
			Comments:       u.Comments,
			UpdatedAt:      u.UpdatedAt,
		}
	}

	return Response{
		Items: users,
	}
}

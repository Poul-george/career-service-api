package create

import (
	"context"
	"fmt"
	"github.com/Poul-george/go-api/api/util/errors"

	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/core/domain/repository/user"
)

type UseCase struct {
	UserRepository user.Repository
}

func NewUseCase(
	userRepository user.Repository,
) *UseCase {
	return &UseCase{
		UserRepository: userRepository,
	}
}

func (u *UseCase) Do(ctx context.Context, ip Input) error {
	res, err := InputData(ip)
	if err != nil {
		return errors.Wrap(err)
	}

	user, err := model.NewUser(res.ExternalUserID, res.Name, res.Password, res.MailAddress, res.Comments)
	if err != nil {
		return errors.Wrap(err)
	}

	fmt.Printf("--------------------{[%v] user create}----------------------------00000\n", u)

	err = u.UserRepository.Create(ctx, user)
	if err != nil {
		return errors.Wrap(err)
	}

	return nil
}

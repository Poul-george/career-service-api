package detail

import (
	"context"
	"fmt"
	"github.com/Poul-george/go-api/api/util/errors"

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

func (u *UseCase) Do(ctx context.Context, input Input) (*Output, error) {
	user, err := u.UserRepository.FindByID(ctx, input.ExternalUserID, input.UserID)
	fmt.Printf("--------------------{[%v]} user detail----------------------------\n", user)

	if err != nil {
		return nil, errors.Wrap(err)
	}

	return NewOutput(*user), nil
}

package detail

import (
	"context"
	"github.com/Poul-george/go-api/api/util/errors"

	"github.com/Poul-george/go-api/api/core/domain/repository/user"
)

type UseCase struct {
	userRepository user.Repository
}

func NewUseCase(
	userRepository user.Repository,
) *UseCase {
	return &UseCase{
		userRepository: userRepository,
	}
}

func (u *UseCase) Do(ctx context.Context, input Input) (*Output, error) {
	user, err := u.userRepository.FindByID(ctx, input.ExternalUserID, input.UserID)

	if err != nil {
		return nil, errors.Wrap(err)
	}

	return NewOutput(*user), nil
}

package list

import (
	"context"
	"github.com/Poul-george/go-api/api/util/errors"

	// "github.com/Poul-george/go-api/api/infrastructure/repository/user"
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

func (u *UseCase) Do(ctx context.Context) (*Output, error) {
	users, err := u.UserRepository.FindByIDs(ctx)

	if err != nil {
		return nil, errors.Wrap(err)
	}

	return NewOutput(users), nil
}

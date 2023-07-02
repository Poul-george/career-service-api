package list

import (
	"context"
	"fmt"
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
	fmt.Printf("--------------------{[%v]}----------------------------00000\n", users)

	if err != nil {
		return nil, errors.Wrap(err)
	}

	return NewOutput(users), nil
}

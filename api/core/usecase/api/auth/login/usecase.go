package login

import (
	"context"
	"github.com/Poul-george/go-api/api/config"
	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/core/domain/repository/auth/token"
	"github.com/Poul-george/go-api/api/core/domain/repository/user"
	"github.com/Poul-george/go-api/api/util/errors"
)

type UseCase struct {
	userRepository  user.Repository
	tokenRepository token.Repository
	serviceConfig   config.Service
}

func NewUseCase(
	userRepository user.Repository,
	tokenRepository token.Repository,
	serviceConfig config.Service,
) *UseCase {
	return &UseCase{
		userRepository:  userRepository,
		tokenRepository: tokenRepository,
		serviceConfig:   serviceConfig,
	}
}

func (u *UseCase) Do(ctx context.Context, input Input) (*Output, error) {
	user, err := u.userRepository.FindByID(ctx, input.ExternalUserID, input.UserID)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	authToken := model.NewAuthToken(user.ID(), user.ExternalUserID())

	accessToken, err := u.tokenRepository.CreateAccessToken(ctx, *authToken, u.serviceConfig.SecretKey)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	res := newOutput(*authToken, *accessToken)
	return &res, nil
}

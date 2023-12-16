package token

import (
	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/util/errors"
	"github.com/Poul-george/go-api/api/util/localtime"
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

const jwtHeaderToken = "auth_token"

type jwtClaims struct {
	IssuedAt int64     `json:"iat" validate:"required"`
	UserID   uint64    `json:"uid" validate:"required"`
	JwtID    uuid.UUID `json:"jid" validate:"required"`
}

func generateJWT(authToken model.AuthToken) (*string, error) {
	now := localtime.Now()

	claims := jwtClaims{
		IssuedAt: now.Unix(),
		UserID:   authToken.UserID().Uint64(),
		JwtID:    uuid.New(),
	}

	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: []byte("app_key")}, (&jose.SignerOptions{}).WithType(jwtHeaderToken))
	if err != nil {
		return nil, errors.Wrap(err)
	}

	signed, err := jwt.Signed(sig).Claims(claims).CompactSerialize()
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &signed, nil
}

package token

import (
	"github.com/Poul-george/go-api/api/core/common/types/identifier"
	"github.com/Poul-george/go-api/api/core/domain/model"
	"github.com/Poul-george/go-api/api/util/convert"
	"github.com/Poul-george/go-api/api/util/errors"
	"github.com/Poul-george/go-api/api/util/localtime"
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

const jwtHeaderToken = "auth_token"

type jwtClaims struct {
	IssuedAt     int64     `json:"iat" validate:"required"`
	ExpirationAt int64     `json:"eat" validate:"required"`
	UserID       uint64    `json:"uid" validate:"required"`
	JwtID        uuid.UUID `json:"jid" validate:"required"`
}

func generateJWT(authToken model.AuthToken, appKey string) (*string, error) {
	now := localtime.Now()

	claims := jwtClaims{
		IssuedAt:     now.Unix(),
		ExpirationAt: int64(30 * time.Minute), // 30分
		UserID:       authToken.UserID().Uint64(),
		JwtID:        uuid.New(),
	}

	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: []byte(appKey)}, (&jose.SignerOptions{}).WithType(jwtHeaderToken))
	if err != nil {
		return nil, errors.Wrap(err)
	}

	signed, err := jwt.Signed(sig).Claims(claims).CompactSerialize()
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return &signed, nil
}

func decodeJWT(signed, appKey string) (*identifier.UserID, error) {
	token, err := jwt.ParseSigned(signed)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	claims := jwtClaims{}
	if err := token.Claims([]byte(appKey), &claims); err != nil {
		return nil, errors.Wrap(err)
	}

	if localtime.Now().After(time.Unix(claims.ExpirationAt, 0)) {
		return nil, errors.New("token expired")
	}

	return convert.ToP(identifier.UserID(claims.UserID)), nil
}

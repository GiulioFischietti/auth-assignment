package crypto

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int64 `json:"user_id"`

	jwt.RegisteredClaims
}

// generates a jwt with this payload (example):
//
//	{
//	 "user_id":10,
//	 "aud":"orders-service",
//	 "exp":1752190000
//	}
func GenerateJWT(
	userID int64,
	service string,
	issuer string,
	privateKey interface{},
) (string, error) {

	claims := Claims{

		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: issuer,

			Subject: strconv.FormatInt(
				userID,
				10,
			),

			Audience: jwt.ClaimStrings{
				service,
			},

			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(
					5 * time.Minute,
				),
			),

			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		claims,
	)

	return token.SignedString(privateKey)
}

package crypto_test

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
	"time"

	"auth-service/crypto"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateJWT(t *testing.T) {

	// temp RSA key
	privateKey, err :=
		rsa.GenerateKey(
			rand.Reader,
			2048,
		)

	if err != nil {
		t.Fatal(err)
	}

	tokenString, err :=
		crypto.GenerateJWT(
			123,
			"orders-service",
			"auth-service",
			privateKey,
		)

	if err != nil {

		t.Fatalf(
			"failed generating token: %v",
			err,
		)
	}

	if tokenString == "" {

		t.Fatal(
			"token should not be empty",
		)
	}

	// check public key jwt

	token, err :=
		jwt.Parse(
			tokenString,

			func(
				token *jwt.Token,
			) (interface{}, error) {

				// check algorithm field claim

				if token.Method.Alg() != "RS256" {

					t.Fatalf(
						"wrong algorithm: %s",
						token.Method.Alg(),
					)
				}

				return &privateKey.PublicKey, nil
			},
		)

	if err != nil {

		t.Fatalf(
			"token validation failed: %v",
			err,
		)
	}

	if !token.Valid {

		t.Fatal(
			"token is not valid",
		)
	}

	claims, ok :=
		token.Claims.(jwt.MapClaims)

	if !ok {

		t.Fatal(
			"invalid claims",
		)
	}

	// check subject field claim

	if claims["sub"] != "123" {

		t.Fatalf(
			"expected sub 123, got %v",
			claims["sub"],
		)
	}

	// check issuer field claim

	if claims["iss"] != "auth-service" {

		t.Fatalf(
			"wrong issuer: %v",
			claims["iss"],
		)
	}

	// check audience field claim

	aud, ok :=
		claims["aud"].([]interface{})

	if !ok {

		t.Fatal(
			"aud claim missing",
		)
	}

	if aud[0] != "orders-service" {

		t.Fatalf(
			"wrong audience: %v",
			aud[0],
		)
	}

	// check expiration field claim

	exp, ok :=
		claims["exp"].(float64)

	if !ok {

		t.Fatal(
			"expiration missing",
		)
	}

	expTime :=
		time.Unix(
			int64(exp),
			0,
		)

	if expTime.Before(
		time.Now(),
	) {

		t.Fatal(
			"token already expired",
		)
	}
}

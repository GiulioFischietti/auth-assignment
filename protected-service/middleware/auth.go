package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func JWTMiddleware(
	publicKey interface{},
	issuer string,
	audience string,
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(
		func(
			w http.ResponseWriter,
			r *http.Request,
		) {

			tokenString :=
				r.Header.Get(
					"Authorization",
				)

			if tokenString == "" {

				http.Error(
					w,
					"missing token",
					http.StatusUnauthorized,
				)

				return
			}

			token, err :=
				jwt.Parse(
					tokenString,

					func(
						token *jwt.Token,
					) (interface{}, error) {

						if token.Method.Alg() != "RS256" {
							return nil,
								jwt.ErrSignatureInvalid
						}

						return publicKey, nil
					},

					jwt.WithIssuer(
						issuer,
					),

					jwt.WithAudience(
						audience,
					),
				)

			if err != nil || !token.Valid {

				http.Error(
					w,
					"invalid token",
					http.StatusUnauthorized,
				)

				return
			}

			claims, ok :=
				token.Claims.(jwt.MapClaims)

			if !ok {

				http.Error(
					w,
					"invalid claims",
					401,
				)

				return
			}

			userID, ok :=
				claims["sub"].(string)

			if !ok {
				http.Error(
					w,
					"missing subject",
					401,
				)
				return
			}

			ctx :=
				context.WithValue(
					r.Context(),
					UserIDKey,
					userID,
				)

			next.ServeHTTP(
				w,
				r.WithContext(ctx),
			)
		},
	)
}

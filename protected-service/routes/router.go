package routes

import (
	"net/http"

	"protected-service/config"
	"protected-service/handlers"
	"protected-service/middleware"
)

func SetupRouter(
	orderHandler *handlers.OrderHandler,
	publicKey interface{},
	cfg config.Config,
) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/health",
		handlers.Health,
	)

	protectedOrders :=
		middleware.JWTMiddleware(
			publicKey,
			cfg.JWTIssuer,
			cfg.JWTAudience,
			http.HandlerFunc(
				orderHandler.GetOrders,
			),
		)

	mux.Handle(
		"/orders",
		protectedOrders,
	)

	return mux
}

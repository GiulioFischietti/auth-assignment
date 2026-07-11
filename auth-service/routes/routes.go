package routes

import (
	"net/http"

	"auth-service/handlers"
)

func SetupRouter(
	authHandler *handlers.AuthHandler,
) http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/health",
		handlers.Health,
	)

	mux.HandleFunc(
		"/register",
		authHandler.Register,
	)

	mux.HandleFunc(
		"/login",
		authHandler.Login,
	)

	mux.HandleFunc(
		"/token",
		authHandler.Token,
	)

	mux.HandleFunc(
		"/logout",
		authHandler.Logout,
	)

	return mux
}

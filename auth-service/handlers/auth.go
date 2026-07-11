package handlers

import (
	"encoding/json"
	"net/http"

	"auth-service/services"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(
	authService *services.AuthService,
) *AuthHandler {

	return &AuthHandler{
		authService: authService,
	}
}

// POST /register
func (h *AuthHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req RegisterRequest

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"invalid body",
			http.StatusBadRequest,
		)

		return
	}

	err = h.authService.Register(
		r.Context(),
		req.Username,
		req.Password,
	)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.WriteHeader(
		http.StatusCreated,
	)

	w.Write(
		[]byte(`{"message":"user created"}`),
	)
}

// POST /login
func (h *AuthHandler) Login(
	w http.ResponseWriter,
	r *http.Request,
) {

	var req LoginRequest

	err := json.NewDecoder(
		r.Body,
	).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"invalid body",
			400,
		)

		return
	}

	token, err :=
		h.authService.Login(
			r.Context(),
			req.Username,
			req.Password,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			401,
		)
		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		map[string]string{
			"session_token": token,
		},
	)
}

// POST /token
func (h *AuthHandler) Token(
	w http.ResponseWriter,
	r *http.Request,
) {

	sessionToken :=
		r.Header.Get(
			"Authorization",
		)

	if sessionToken == "" {

		http.Error(
			w,
			"missing session token",
			401,
		)

		return
	}

	var req TokenRequest

	err :=
		json.NewDecoder(
			r.Body,
		).Decode(&req)

	if err != nil {

		http.Error(
			w,
			"invalid body",
			400,
		)

		return
	}

	accessToken, err :=
		h.authService.CreateAccessToken(
			r.Context(),
			sessionToken,
			req.ServiceName,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			401,
		)

		return
	}

	json.NewEncoder(w).Encode(
		map[string]string{
			"access_token": accessToken,
		},
	)
}

// POST /logout
func (h *AuthHandler) Logout(
	w http.ResponseWriter,
	r *http.Request,
) {
	sessionToken :=
		r.Header.Get("Authorization")

	if sessionToken == "" {
		http.Error(
			w,
			"missing token",
			http.StatusUnauthorized,
		)
		return
	}
	err :=
		h.authService.Logout(
			r.Context(),
			sessionToken,
		)

	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(
		http.StatusNoContent,
	)
}

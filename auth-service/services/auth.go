package services

import (
	"context"
	"errors"
	"time"

	"auth-service/crypto"
	"auth-service/models"
	"auth-service/repositories"
)

type AuthService struct {
	users      repositories.UserRepository
	sessions   repositories.SessionRepository
	services   repositories.ServiceRepository
	issuer     string
	privateKey interface{}
}

func NewAuthService(
	users repositories.UserRepository,
	sessions repositories.SessionRepository,
	services repositories.ServiceRepository,
	issuer string,
	privateKey interface{},
) *AuthService {

	return &AuthService{
		users:      users,
		sessions:   sessions,
		services:   services,
		privateKey: privateKey,
		issuer:     issuer,
	}
}

// Login:
// username/password -> session token
func (a *AuthService) Login(
	ctx context.Context,
	username string,
	password string,
) (string, error) {

	user, err := a.users.FindByUsername(
		ctx,
		username,
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !crypto.CheckPassword(
		password,
		user.PasswordHash,
	) {
		return "", errors.New("invalid credentials")
	}

	rawToken, err := crypto.GenerateToken()

	if err != nil {
		return "", err
	}

	session := &models.Session{
		UserID: user.ID,

		SessionTokenHash: crypto.HashToken(rawToken),

		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	err = a.sessions.Create(
		ctx,
		session,
	)

	if err != nil {
		return "", err
	}

	return rawToken, nil
}

// Validate session token
func (a *AuthService) ValidateSession(
	ctx context.Context,
	rawToken string,
) (*models.Session, error) {

	hash := crypto.HashToken(rawToken)

	session, err :=
		a.sessions.FindByTokenHash(
			ctx,
			hash,
		)

	if err != nil {
		return nil, errors.New("invalid session")
	}

	// if session is cached, we don't need to check revoked field.
	// In redis, if a session token has been revoked, it is simply removed
	if session.Cached {
		return session, nil
	}

	if session.RevokedAt != nil {
		return nil, errors.New("session revoked")
	}

	if time.Now().After(
		session.ExpiresAt,
	) {
		return nil, errors.New("session expired")
	}

	return session, nil
}

// Generate JWT access token for a service
func (a *AuthService) CreateAccessToken(
	ctx context.Context,
	sessionToken string,
	serviceName string,
) (string, error) {

	session, err :=
		a.ValidateSession(
			ctx,
			sessionToken,
		)

	if err != nil {
		return "", err
	}

	service, err :=
		a.services.FindByName(
			ctx,
			serviceName,
		)

	if err != nil {
		return "", errors.New("service not found")
	}

	if !service.Active {
		return "", errors.New("service inactive")
	}

	return crypto.GenerateJWT(
		session.UserID,
		service.Name,
		a.issuer,
		a.privateKey,
	)
}

func (a *AuthService) Register(
	ctx context.Context,
	username string,
	password string,
) error {

	passwordHash, err := crypto.HashPassword(password)

	if err != nil {
		return err
	}

	user := &models.User{
		Username:     username,
		PasswordHash: passwordHash,
	}

	err = a.users.Create(
		ctx,
		user,
	)

	if err != nil {
		return errors.New("your request could not be processed")
	}

	return nil
}

func (a *AuthService) Logout(
	ctx context.Context,
	sessionToken string,
) error {

	hash := crypto.HashToken(
		sessionToken,
	)

	return a.sessions.Revoke(
		ctx,
		hash,
	)
}

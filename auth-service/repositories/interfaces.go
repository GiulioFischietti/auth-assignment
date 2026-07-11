package repositories

import (
	"context"

	"auth-service/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByUsername(ctx context.Context, username string) (*models.User, error)
}

type SessionRepository interface {
	Create(ctx context.Context, session *models.Session) error
	FindByTokenHash(ctx context.Context, hash string) (*models.Session, error)
	Revoke(ctx context.Context, hash string) error
}

type ServiceRepository interface {
	FindByName(ctx context.Context, name string) (*models.Service, error)
}

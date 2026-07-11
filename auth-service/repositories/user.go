package repositories

import (
	"context"

	"auth-service/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(
	ctx context.Context,
	user *models.User,
) error {

	query := `
		INSERT INTO users(username, password_hash)
		VALUES($1,$2)
		RETURNING id
	`

	return r.db.QueryRow(
		ctx,
		query,
		user.Username,
		user.PasswordHash,
	).Scan(&user.ID)
}

func (r *userRepository) FindByUsername(
	ctx context.Context,
	username string,
) (*models.User, error) {

	query := `
		SELECT id, username, password_hash
		FROM users
		WHERE username=$1
	`

	user := &models.User{}

	err := r.db.QueryRow(
		ctx,
		query,
		username,
	).Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

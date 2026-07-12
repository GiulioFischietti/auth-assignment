package repositories

import (
	"context"

	"auth-service/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type persistedSessionRepository struct {
	db *pgxpool.Pool
}

func NewPersistedSessionRepository(db *pgxpool.Pool) SessionRepository {
	return &persistedSessionRepository{
		db: db,
	}
}

func (r *persistedSessionRepository) Create(
	ctx context.Context,
	session *models.Session,
) error {

	query := `
		INSERT INTO sessions(
			user_id,
			session_token_hash,
			expires_at
		)
		VALUES($1,$2,$3)
		RETURNING id
	`

	return r.db.QueryRow(
		ctx,
		query,
		session.UserID,
		session.SessionTokenHash,
		session.ExpiresAt,
	).Scan(&session.ID)
}

func (r *persistedSessionRepository) FindByTokenHash(
	ctx context.Context,
	hash string,
) (*models.Session, error) {

	query := `
		SELECT 
			id,
			user_id,
			session_token_hash,
			expires_at,
			revoked_at
		FROM sessions
		WHERE session_token_hash=$1
	`

	session := &models.Session{}

	err := r.db.QueryRow(
		ctx,
		query,
		hash,
	).Scan(
		&session.ID,
		&session.UserID,
		&session.SessionTokenHash,
		&session.ExpiresAt,
		&session.RevokedAt,
	)

	if err != nil {
		return nil, err
	}

	return session, nil
}

func (r *persistedSessionRepository) Revoke(
	ctx context.Context,
	hash string,
) error {

	_, err := r.db.Exec(
		ctx,
		`
		UPDATE sessions
		SET revoked_at = NOW()
		WHERE session_token_hash = $1
		`,
		hash,
	)

	return err
}

package repositories

import (
	"context"

	"auth-service/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type serviceRepository struct {
	db *pgxpool.Pool
}

func NewServiceRepository(
	db *pgxpool.Pool,
) ServiceRepository {

	return &serviceRepository{
		db: db,
	}
}

func (r *serviceRepository) FindByName(
	ctx context.Context,
	name string,
) (*models.Service, error) {

	query := `
		SELECT id,name,active
		FROM services
		WHERE name=$1
	`

	service := &models.Service{}

	err := r.db.QueryRow(
		ctx,
		query,
		name,
	).Scan(
		&service.ID,
		&service.Name,
		&service.Active,
	)

	if err != nil {
		return nil, err
	}

	return service, nil
}

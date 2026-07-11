package database

import (
	"context"
	"fmt"

	"auth-service/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(cfg *config.Config) (*pgxpool.Pool, error) {

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := pgxpool.New(
		context.Background(),
		connectionString,
	)

	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())

	if err != nil {
		return nil, err
	}

	return db, nil
}

package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	MongoURI      string
	MongoDatabase string
	PublicKeyPath string
	JWTIssuer     string
	JWTAudience   string
}

func Load() *Config {

	_ = godotenv.Load()

	return &Config{

		Port: getEnv(
			"PORT",
			"8081",
		),

		MongoURI: getEnv(
			"MONGO_URI",
			"mongodb://localhost:27017",
		),

		MongoDatabase: getEnv(
			"MONGO_DATABASE",
			"orders_db",
		),

		PublicKeyPath: getEnv(
			"PUBLIC_KEY_PATH",
			"./keys/public.pem",
		),

		JWTIssuer: getEnv(
			"JWT_ISSUER",
			"auth-service",
		),

		JWTAudience: getEnv(
			"JWT_AUDIENCE",
			"orders-service",
		),
	}
}

func getEnv(
	key string,
	fallback string,
) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

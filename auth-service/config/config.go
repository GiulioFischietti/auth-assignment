package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	PrivateKeyPath string
	PublicKeyPath  string

	JWTIssuer string
}

var App *Config

func Load() *Config {

	// Ignore error if .env does not exist (e.g. Docker)
	_ = godotenv.Load()

	App = &Config{
		Port: getEnv("PORT", "8080"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres"),
		DBName:     getEnv("DB_NAME", "auth_db"),

		PrivateKeyPath: getEnv("PRIVATE_KEY_PATH", "./keys/private.pem"),
		PublicKeyPath:  getEnv("PUBLIC_KEY_PATH", "./keys/public.pem"),

		JWTIssuer: getEnv("JWT_ISSUER", "auth-service"),
	}

	log.Println("Configuration loaded")

	return App
}

func getEnv(key string, fallback string) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

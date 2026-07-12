package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int

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

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvInt("REDIS_DB", 0),

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

func getEnvInt(key string, fallback int) int {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	number, err := strconv.Atoi(value)

	if err != nil {
		return fallback
	}

	return number
}

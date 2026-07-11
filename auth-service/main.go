package main

import (
	"log"
	"net/http"

	"auth-service/config"
	"auth-service/crypto"
	"auth-service/database"
	"auth-service/handlers"
	"auth-service/repositories"
	"auth-service/routes"
	"auth-service/services"
)

func main() {

	cfg := config.Load()

	privateKey, err := crypto.LoadPrivateKey(cfg.PrivateKeyPath)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("RSA private key loaded daje")
	}

	db, err := database.Connect(cfg)

	if err != nil {
		log.Fatal("database connection failed:", err)
	} else {
		log.Print("database connected daje")
	}

	defer db.Close()

	userRepo := repositories.NewUserRepository(db)

	sessionRepo := repositories.NewSessionRepository(db)

	serviceRepo := repositories.NewServiceRepository(db)

	authService := services.NewAuthService(
		userRepo,
		sessionRepo,
		serviceRepo,
		cfg.JWTIssuer,
		privateKey,
	)

	authHandler := handlers.NewAuthHandler(authService)

	router := routes.SetupRouter(authHandler)

	log.Println("server running on port", cfg.Port)

	err = http.ListenAndServe(
		":"+cfg.Port,
		router,
	)

	if err != nil {
		log.Fatal(err)
	}
}

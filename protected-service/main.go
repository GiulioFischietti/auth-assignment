package main

import (
	"context"
	"log"
	"net/http"

	"protected-service/config"
	"protected-service/crypto"
	"protected-service/database"
	"protected-service/handlers"
	"protected-service/repositories"
	"protected-service/routes"
	"protected-service/services"
)

func main() {

	cfg :=
		config.Load()

	publicKey, err :=
		crypto.LoadPublicKey(
			cfg.PublicKeyPath,
		)

	if err != nil {

		log.Fatal(
			"cannot load public key:",
			err,
		)
	}

	mongoClient, err :=
		database.ConnectMongo(
			cfg.MongoURI,
		)

	if err != nil {

		log.Fatal(
			"mongo connection failed:",
			err,
		)
	}

	defer mongoClient.Disconnect(
		context.Background(),
	)

	db :=
		mongoClient.Database(
			cfg.MongoDatabase,
		)

	orderRepository :=
		repositories.NewOrderRepository(
			db,
		)

	orderService :=
		services.NewOrderService(
			orderRepository,
		)

	orderHandler :=
		handlers.NewOrderHandler(
			orderService,
		)

	router :=
		routes.SetupRouter(
			orderHandler,
			publicKey,
			*cfg,
		)

	log.Println(
		"protected service running on",
		cfg.Port,
	)

	err =
		http.ListenAndServe(
			":"+cfg.Port,
			router,
		)

	if err != nil {

		log.Fatal(err)
	}
}

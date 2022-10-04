package app

import (
	"backend-face/internal/config"
	"backend-face/internal/controller"
	"backend-face/internal/repository"
	"backend-face/internal/usecase"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func Run(cfg config.Config) {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(cfg.MongodbURL)
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatalln(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalln(err)
	}

	collection := client.Database(cfg.MongodbDatabaseName).Collection(cfg.MongodbUsersCollection)

	userRepository := repository.NewUserRepository(*collection)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	userController.RegisterHandlers()

	log.Printf("Starting HTTP server on port %s\n", cfg.HttpPort)

	port := fmt.Sprintf(":%s", cfg.HttpPort)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalln(err)
	}
}

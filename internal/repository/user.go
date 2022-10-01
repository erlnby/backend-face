package repository

import (
	"backend-face/internal/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type UserRepository struct {
	collection mongo.Collection
}

func NewUserRepository(collection mongo.Collection) UserRepository {
	return UserRepository{collection}
}

func (repository UserRepository) GetAll() []entity.User {
	ctx := context.TODO()
	cursor, err := repository.collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatalln(err)
	}
	var users []entity.User
	if err = cursor.All(ctx, &users); err != nil {
		log.Fatalln(err)
	}
	if err = cursor.Close(ctx); err != nil {
		log.Fatalln(err)
	}
	return users
}

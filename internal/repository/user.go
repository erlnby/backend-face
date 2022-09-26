package repository

import (
	"backend-face/internal/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection mongo.Collection
}

func NewUserRepository(collection mongo.Collection) UserRepository {
	return UserRepository{collection}
}

func (repository UserRepository) GetAll() []entity.User {
	/*
		Здесь нужно реализовать выбор юзеров из MongoDB.
	*/
	return []entity.User{}
}

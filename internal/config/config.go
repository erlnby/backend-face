package config

import (
	"fmt"
	"os"
)

type Config struct {
	HttpPort               string
	MongodbURL             string
	MongodbDatabaseName    string
	MongodbUsersCollection string
}

func NewConfig() Config {
	return Config{
		HttpPort: os.Getenv("HTTP_PORT"),
		MongodbURL: fmt.Sprintf("mongodb://%s:%s@db:27017",
			os.Getenv("MONGO_USER"),
			os.Getenv("MONGO_PASSWORD")),
		MongodbDatabaseName:    os.Getenv("MONGODB_DATABASE_NAME"),
		MongodbUsersCollection: os.Getenv("MONGODB_USERS_COLLECTION"),
	}
}

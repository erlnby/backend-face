package config

import "os"

type Config struct {
	HttpPort               string
	MongodbURL             string
	MongodbDatabaseName    string
	MongodbUsersCollection string
}

func NewConfig() Config {
	return Config{
		HttpPort:               os.Getenv("HTTP_PORT"),
		MongodbURL:             os.Getenv("MONGODB_URL"),
		MongodbDatabaseName:    os.Getenv("MONGODB_DATABASE_NAME"),
		MongodbUsersCollection: os.Getenv("MONGODB_USERS_COLLECTION"),
	}
}

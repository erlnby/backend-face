package app

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func init() {

	databaseURL := fmt.Sprintf("mongodb://%s:%s@mongodb-service:27017/%s?authSource=admin",
		os.Getenv("MONGO_USER"),
		os.Getenv("MONGO_PASSWORD"),
		os.Getenv("MONGODB_DATABASE_NAME"))

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	for attempts > 0 {
		m, err = migrate.New("file://migrations", databaseURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: mondodb is trying to connect, attempts left: %d\n", attempts)
		log.Println(err)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate: mondodb connect error: %s\n", err)
	}

	err = m.Up()
	defer func(m *migrate.Migrate) {
		err, _ := m.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(m)
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s\n", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Println("Migrate: no change")
		return
	}

	log.Println("Migrate: up success")
}

package main

import (
	"backend-face/internal/app"
	"backend-face/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env.example")
	_ = godotenv.Load(".env")

	cfg := config.NewConfig()

	app.Run(cfg)
}

package main

import (
	"backend-face/internal/app"
	"backend-face/internal/config"
)

func main() {
	cfg := config.NewConfig()

	app.Run(cfg)
}

package main

import (
	"backend-face/config"
	"backend-face/internal/app"
)

func main() {
	cfg := config.NewConfig()
	app.Run(cfg)
}

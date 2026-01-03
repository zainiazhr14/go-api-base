package main

import (
	"log"

	"github.com/zainiazhr14/go-api/api"
	"github.com/zainiazhr14/go-api/config"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	app := &api.App{}
	app.Initialize(&config)

	log.Printf("Starting %s on port %s\n", config.AppName, config.Port)

	app.Run(":" + config.Port)
}

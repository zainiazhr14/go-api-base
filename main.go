package main

import (
	"github.com/zainiazhr14/go-api/app"
	"github.com/zainiazhr14/go-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}

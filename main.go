package main

import (
	"log"

	"github.com/hawwwdi/YAML-Api/src/app"
	"github.com/hawwwdi/YAML-Api/src/config"
	"github.com/hawwwdi/YAML-Api/src/renderer"
)

func main() {
	configs := config.LoadConfig("./src/config/config.json")
	r := renderer.NewRenderer(configs.Templates)
	app := app.NewApp(r)
	log.Fatal(app.Start(configs.Port))
}

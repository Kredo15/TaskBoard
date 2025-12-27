package main

import (
	"taskboard/config"
	"taskboard/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	app.Run(cfg)

}

package main

import (
	"log"

	di "github.com/Anandhu4456/go-Ecommerce/pkg/di"

	"github.com/Anandhu4456/go-Ecommerce/pkg/config"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config", configErr)
	}
	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("couldn't start server", err)
	} else {
		server.Start()
	}
}

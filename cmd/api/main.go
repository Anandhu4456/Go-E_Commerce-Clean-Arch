package main

import (
	"log"

	"github.com/Anandhu4456/go-Ecommerce/cmd/api/docs"
	di "github.com/Anandhu4456/go-Ecommerce/pkg/di"

	"github.com/Anandhu4456/go-Ecommerce/pkg/config"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config", configErr)
	}

	docs.SwaggerInfo.Title = "YoursStore"
	docs.SwaggerInfo.Description = "An Ecommerce Application"
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = config.BASE_URL
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal("couldn't start server", err)
	} else {
		server.Start()
	}
}

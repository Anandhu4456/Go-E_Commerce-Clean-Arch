package main

import (
	"log"

	"github.com/Anandhu4456/go-Ecommerce/pkg/config"
	"github.com/Anandhu4456/go-Ecommerce/pkg/di"
)

func main(){
	config,configErr:=config.LoadConfig()
	if configErr!=nil{
		log.Fatal("cannot load config",configErr)
	}
	server,err:=di.InitializeAPI(config)
	if err!=nil{
		log.Fatal("couldn't start server",err)
	}else{
		server.Start()
	}
}
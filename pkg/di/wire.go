//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/api"
	
	handler "github.com/Anandhu4456/go-Ecommerce/pkg/api/handlers"
	config "github.com/Anandhu4456/go-Ecommerce/pkg/config"
	db "github.com/Anandhu4456/go-Ecommerce/pkg/db"
	repository "github.com/Anandhu4456/go-Ecommerce/pkg/repository"
	usecase "github.com/Anandhu4456/go-Ecommerce/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.ConnectDB, repository.NewUserRepository, usecase.NewUserUsecase, handler.NewUserHandler,api.NewServerHttp)
	
	return &api.ServerHTTP{},nil
}

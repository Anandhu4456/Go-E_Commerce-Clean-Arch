package di

import (
	http "github.com/Anandhu4456/go-Ecommerce/pkg/api"
	handler "github.com/Anandhu4456/go-Ecommerce/pkg/api/handlers"
	config "github.com/Anandhu4456/go-Ecommerce/pkg/config"
	db "github.com/Anandhu4456/go-Ecommerce/pkg/db"
	repository "github.com/Anandhu4456/go-Ecommerce/pkg/repository"
	usecase "github.com/Anandhu4456/go-Ecommerce/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDB, repository.NewUserRepository, usecase.NewUserUsecase, handler.NewUserHandler, http.NewSeverHTTP)
	return &http.ServerHTTP{}, nil
}

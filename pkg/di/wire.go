package di

import (
	"github.com/Anandhu4456/go-Ecommerce/pkg/api"
	"github.com/Anandhu4456/go-Ecommerce/pkg/api/handlers"
	"github.com/Anandhu4456/go-Ecommerce/pkg/config"
	"github.com/Anandhu4456/go-Ecommerce/pkg/db"
	"github.com/Anandhu4456/go-Ecommerce/pkg/repository"
	"github.com/Anandhu4456/go-Ecommerce/pkg/usecase"
	"github.com/google/wire"
)

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.ConnectDB,
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		handlers.NewUserHandler,
		api.NewServerHttp)
	return &api.ServerHTTP{}, nil
}

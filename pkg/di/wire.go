//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/profile/service/pkg/api"
	"github.com/profile/service/pkg/api/handler"
	"github.com/profile/service/pkg/config"
	"github.com/profile/service/pkg/db"
	"github.com/profile/service/pkg/repository"
	"github.com/profile/service/pkg/usecase"
)

func InitApi(cfg config.Config) (*api.ServerHttp, error) {
	wire.Build(
		db.ConnectToDataBase,
		repository.NewUserRepo,
		usecase.NewUserUseCase,
		handler.NewUserHandler,
		api.NewServerHttp)
	return &api.ServerHttp{}, nil
}

//go run github.com/google/wire/cmd/wire

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/profile/service/pkg/api"
	"github.com/profile/service/pkg/api/handler"
	"github.com/profile/service/pkg/config"
	"github.com/profile/service/pkg/db"
	"github.com/profile/service/pkg/repository"
	"github.com/profile/service/pkg/usecase"
)

// Injectors from wire.go:

func InitApi(cfg config.Config) (*api.ServerHttp, error) {
	gormDB, err := db.ConnectToDataBase(cfg)
	if err != nil {
		return nil, err
	}
	userRepo := repository.NewUserRepo(gormDB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)
	serverHttp := api.NewServerHttp(userHandler)
	return serverHttp, nil
}

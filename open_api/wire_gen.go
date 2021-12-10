// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"net/http"
	"open_api/app"
	"open_api/controller"
	"open_api/middleware"
	"open_api/repository"
	"open_api/service"
)

// Injectors from injector.go:

func InitializeServer() *http.Server {
	activityRepositoryImpl := repository.NewActivityRepository()
	db := app.NewDB()
	validate := validator.New()
	activityService := service.NewActivityService(activityRepositoryImpl, db, validate)
	activityController := controller.NewActivityController(activityService)
	router := app.NewRouter(activityController)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var categoryRepoSet = wire.NewSet(repository.NewActivityRepository, wire.Bind(new(repository.ActivityRepository), new(*repository.ActivityRepositoryImpl)))
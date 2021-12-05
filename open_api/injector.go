//go:build wireinject
//+build wireinject

package main

import (
	"net/http"
	"open_api/app"
	"open_api/controller"
	"open_api/repository"
	"open_api/service"
	"open_api/middleware"

	"github.com/julienschmidt/httprouter"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
)

var categoryRepoSet = wire.NewSet(
	repository.NewActivityRepository,
	wire.Bind(new(repository.ActivityRepository), new(*repository.ActivityRepositoryImpl)),
)

func InitializeServer() *http.Server {
	// mydb 				:= app.NewDB()
	// myValidator := validator.New()

	// activityRepository 	:= repository.NewActivityRepository()
	// activityService 		:= service.NewActivityService(activityRepository, mydb, myValidator)
	// activityController 	:= controller.NewActivityController(activityService)

	// router := app.NewRouter(activityController)
	// server := http.Server{
	// 	Addr:    "localhost:3000",
	// 	Handler: middleware.NewAuthMiddleware(router),
	// }

	wire.Build(
		app.NewDB,
		validator.New,
		categoryRepoSet,
		service.NewActivityService,
		controller.NewActivityController,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		// google/wire can NOT handle multi provider with same type parameter, so we make NewServer func on main.go
		// wire.Bind(new(http.Handler), new(*middleware.AuthMiddleware)),
		// wire.Struct(new(http.Server), "Handler"),
		NewServer,
	)
	return nil
}
package main

import (
	"fmt"
	"net/http"
	"open_api/app"
	"open_api/controller"
	"open_api/helper"
	"open_api/middleware"
	"open_api/repository"
	"open_api/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	mydb 				:= app.NewDB()
	myValidator := validator.New()

	activityRepository 	:= repository.NewActivityRepository()
	activityService 		:= service.NewActivityService(activityRepository, mydb, myValidator)
	activityController 	:= controller.NewActivityController(activityService)

	router := app.NewRouter(activityController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("Server ran on port 3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

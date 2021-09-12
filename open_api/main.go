package main

import (
	"fmt"
	"net/http"
	"open_api/app"
	"open_api/controller"
	"open_api/exception"
	"open_api/helper"
	"open_api/middleware"
	"open_api/repository"
	"open_api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	mydb := app.NewDB()
	myValidator := validator.New()

	activityRepository := repository.NewActivityRepository()
	activityService := service.NewActivityService(activityRepository, mydb, myValidator)
	activityController := controller.NewActivityController(activityService)

	router := httprouter.New()
	router.POST("/api/activities", activityController.Create)
	router.PUT("/api/activities/:activityId", activityController.Update)
	router.DELETE("/api/activities/:activityId", activityController.Delete)

	router.GET("/api/activities", activityController.FindAll)
	router.GET("/api/activities/:activityId", activityController.FindById)

	fmt.Println("Impressive, server running now...!")

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

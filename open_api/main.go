package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"open_api/repository"
	"open_api/service"
	"open_api/controller"
	"open_api/app"
	"open_api/helper"
)

func main() {
	mydb 				:= app.NewDB()
	myValidator := validator.New()

	activityRepository 	:= repository.NewActivityRepository()
	activityService 		:= service.NewActivityService(activityRepository, mydb, myValidator)
	activityController 	:= controller.NewActivityController(activityService)

	router := httprouter.New()
	router.POST("/api/activities", 									activityController.Create)
	router.PUT("/api/activities/:activitiesId", 		activityController.Update)
	router.DELETE("/api/activities/:activitiesId", 	activityController.Delete)

	router.GET("/api/activities", 									activityController.FindAll)
	router.GET("/api/activities/:activitiesId", 		activityController.FindById)

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

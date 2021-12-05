package main

import (
	"fmt"
	"net/http"
	"open_api/helper"
	"open_api/middleware"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

// go run main.go wire_gen.go
func main() {
	// mydb 				:= app.NewDB()
	// myValidator := validator.New()

	// activityRepository	:= repository.NewActivityRepository()
	// activityService			:= service.NewActivityService(activityRepository, mydb, myValidator)
	// activityController	:= controller.NewActivityController(activityService)

	// router := app.NewRouter(activityController)
	// server := http.Server{
	// 	Addr:    "localhost:3000",
	// 	Handler: middleware.NewAuthMiddleware(router),
	// }
	server := InitializeServer()

	fmt.Println("Server ran on port 3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

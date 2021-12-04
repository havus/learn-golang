package app

import (
	"open_api/controller"
	"open_api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(activityController controller.ActivityController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/api/activities", activityController.Create)
	router.PUT("/api/activities/:activityId", activityController.Update)
	router.DELETE("/api/activities/:activityId", activityController.Delete)

	router.GET("/api/activities", activityController.FindAll)
	router.GET("/api/activities/:activityId", activityController.FindById)

	router.PanicHandler = exception.ErrorHandler

	return router
}

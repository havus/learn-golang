package main

import (
	"fmt"
	"learn-gin/book"
	"learn-gin/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// LAYERS
	// main.go
	// handler/controller
	// service
	// repository model
	// db
	// mysql driver

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
  dsn := "root:rootroot@tcp(127.0.0.1:3306)/learn_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&book.Book{})

	// bookFileRepo	:= book.NewRepoFile()
	bookRepo 			:= book.NewRepo(db)
	bookService 	:= book.NewService(bookRepo)
	// bookService 	:= book.NewService(bookFileRepo)
	bookHandler 	:= handler.NewBookHandler(bookService)
	
	router := gin.Default()

	// router.GET("/", rootHandler)
	// router.GET("/books/:id", handler.GetBookHandler)
	
	// VERSIONING
	api := router.Group("/api")
	v1 	:= api.Group("/v1")

	v1.GET("/books", bookHandler.FindAll)
	v1.GET("/books/:id", bookHandler.GetBookHandlerById)
	v1.POST("/books", bookHandler.Create)
	v1.PUT("/books/:id", bookHandler.Update)
	v1.DELETE("/books/:id", bookHandler.Delete)

	router.Run(":3000")
}

// func rootHandler (c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"name": "John Doe",
// 		"bio": 	"learn faster than the others",
// 	})
// }

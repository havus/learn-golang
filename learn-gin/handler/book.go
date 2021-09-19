package handler

import (
	"fmt"
	"learn-gin/book"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *BookHandler {
	return &BookHandler{bookService}
}

// localhost:3000/books/90?name=sherlock holmes
// learn url params, query params
// func GetBookHandler(c *gin.Context) {
// 	id 		:= c.Param("id")
// 	name 	:= c.Query("name")

// 	c.JSON(http.StatusOK, gin.H{
// 		"id": 	id,
// 		"name": name,
// 		"bio": 	"learn faster than the others",
// 	})
// }

func (h *BookHandler) Create(c *gin.Context) {
	var bookParams book.BookRequest

	if err := c.ShouldBindJSON(&bookParams); err != nil {
		// catch "price": "100000sa",
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		
		// don't catch "price": "100000sa",
		
		// errorMessages := []string{}

		// exception, _ := err.(validator.ValidationErrors)
		// fmt.Println(exception)

		// for _, e := range exception {
		// 	errorMessages = append(errorMessages, fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag()))
		// }

		// c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})

		return
	}

	book, _ := h.bookService.Create(bookParams)

	c.JSON(http.StatusCreated, gin.H{
		"data": convertToBookResponse(book),
	})
}

func (h *BookHandler) Update(c *gin.Context) {
	id := c.Param("id")
	intId, _ := strconv.Atoi(id)

	var bookParams book.BookRequest

	if err := c.ShouldBindJSON(&bookParams); err != nil {
		var errorMessages []string

		validationErrors := err.(validator.ValidationErrors)

		for _, e := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag()))
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": errorMessages})
		return
	}

	book, _ := h.bookService.Update(intId, bookParams)

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})
}

func (h *BookHandler) Delete(c *gin.Context) {
	id 				:= c.Param("id")
	intId, _ 	:= strconv.Atoi(id)
	book, _ 	:= h.bookService.Delete(intId)

	c.JSON(http.StatusOK, gin.H{
		"deleted_data": convertToBookResponse(book),
	})
}

func (h *BookHandler) GetBookHandlerById(c *gin.Context) {
	id 				:= c.Param("id")
	intId, _ 	:= strconv.Atoi(id)
	book, err := h.bookService.FindById(intId)

	// https://go.dev/blog/error-handling-and-go
	// fmt.Println(reflect.TypeOf(err)) // *errors.errorString
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": convertToBookResponse(book),
	})
}

func (h *BookHandler) FindAll(c *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var bookResponses []book.BookResponse

	for _, record := range books {
		bookResponses = append(bookResponses, convertToBookResponse(record))
	}

	c.JSON(http.StatusOK, gin.H{
		"data": bookResponses,
	})
}

func convertToBookResponse(record book.Book) book.BookResponse {
	return book.BookResponse{
		Id: 					record.Id,
		Title: 				record.Title,
		Description: 	record.Description,
		Price: 				record.Price,
		Rating: 			record.Rating,
	}
}

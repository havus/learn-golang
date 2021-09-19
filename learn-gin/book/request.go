package book

import "encoding/json"

type BookRequest struct {
	Title 			string			`json:"title" binding:"required"`
	// Price 			int					`json:"price" binding:"required,number"`
	Price 			json.Number	`binding:"required,number"`
	Rating			json.Number	`json:"rating"`
	Description string			`json:"description"`
}
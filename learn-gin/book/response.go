package book

type BookResponse struct {
	Id 					uint 		`json:"id"`
	Title 			string 	`json:"title"`
	Price 			int 		`json:"price"`
	Rating 			int			`json:"rating"`
	Description string	`json:"description"`
}

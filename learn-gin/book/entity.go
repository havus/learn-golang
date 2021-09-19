package book

import "time"

type Book struct {
	// https://docs.microsoft.com/en-us/sql/t-sql/data-types/int-bigint-smallint-and-tinyint-transact-sql?view=sql-server-ver15
	Id 					uint 		`json:"id" gorm:"primaryKey"`
	Title 			string 	`json:"title" gorm:"type:VARCHAR(255) NOT NULL;index"`
	Price 			int // prepare for minus price
	Rating 			int			`json:"rating" gorm:"type: tinyint;default:0"`
	Description string	`json:"description" gorm:"type:VARCHAR(255)"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

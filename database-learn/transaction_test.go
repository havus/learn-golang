package database_learn

import (
	"testing"
	"fmt"
	"context"
	"strconv"
)

func TestTransactionDb(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	query := "INSERT INTO products(name, price) VALUES(?, ?)"

	for i := 11; i <= 20; i ++ {
		productName 	:= "Product #" + strconv.Itoa(i)
		productPrice 	:= i * 10000

		result, err := tx.ExecContext(ctx, query, productName, productPrice)

		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Product ID:", insertId)
	}

	err = tx.Rollback()
	// err = tx.Commit()
	if err != nil {
		panic(err)
	}

	fmt.Println("Pretend something went wrong, rollback transaction!")
}
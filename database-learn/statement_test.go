package database_learn

import (
	"testing"
	"fmt"
	"context"
	"strconv"
)

func TestPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO products(name, price) VALUES(?, ?)"

	statement, err := db.PrepareContext(ctx, query)
	defer statement.Close()

	fmt.Println(statement)

	if err != nil {
		panic(err)
	}

	for i := 1; i <= 10; i ++ {
		productName 	:= "Product #" + strconv.Itoa(i)
		productPrice 	:= i * 10000

		result, err := statement.ExecContext(ctx, productName, productPrice)

		if err != nil {
			panic(err)
		}
		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Product ID:", insertId)
	}
}
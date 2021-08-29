package repository


import (
	database_learn "database-learn"
	"database-learn/entity"
	"testing"
	"context"
	"fmt"
)

func TestInsertProduct(t *testing.T) {
	productRepo := NewProduct(database_learn.GetConnection())
	ctx 				:= context.Background()
	product 		:= entity.Product{
		Name: "Test Repo Product",
		Price: 1000,
		Description: "this is desc",
	}

	result, err := productRepo.Insert(ctx, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindByIdProduct(t *testing.T) {
	productRepo := NewProduct(database_learn.GetConnection())
	ctx 				:= context.Background()

	product, err := productRepo.FindById(ctx, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}

func TestFindAllProduct(t *testing.T) {
	productRepo := NewProduct(database_learn.GetConnection())
	ctx 				:= context.Background()

	products, err := productRepo.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Println(product)
	}
}
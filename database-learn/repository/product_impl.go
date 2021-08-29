package repository

import (
	"database-learn/entity"
	"database/sql"
	"context"
	"errors"
	"strconv"
)

type productRepoImpl struct {
	DB *sql.DB
}

func NewProduct(db *sql.DB) ProductRepository {
	return &productRepoImpl{DB: db}
}

func (repo *productRepoImpl) Insert(ctx context.Context, product entity.Product) (entity.Product, error) {
	query := "INSERT INTO products(name,price,description) VALUES(?, ?, ?)"
	result, err := repo.DB.ExecContext(ctx, query, product.Name, product.Price, product.Description)

	if err != nil {
		return product, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return product, err
	}

	product.Id = id
	return product, nil
}

func (repo *productRepoImpl) FindById(ctx context.Context, id int64) (entity.Product, error) {
	query 		:= "SELECT id, name, price, description FROM products WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, query, id)
	product 	:= entity.Product{}

	defer rows.Close()

	if err != nil {
		return product, err
	}

	if rows.Next() {
		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Description)

		return product, nil
	} else {
		return product, errors.New("Product with id " + strconv.Itoa(int(id)) + " not found!")
	}
}

func (repo *productRepoImpl) FindAll(ctx context.Context) ([]entity.Product, error) {
	query 		:= "SELECT id, name, price, description FROM products"
	rows, err := repo.DB.QueryContext(ctx, query)

	var products []entity.Product

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		product := entity.Product{}

		rows.Scan(&product.Id, &product.Name, &product.Price, &product.Description)

		products = append(products, product)
	}

	return products, nil
}
package repository

import (
	"database-learn/entity"
	"context"
)

type ProductRepository interface {
	Insert(ctx context.Context, product entity.Product) (entity.Product, error)
	FindById(ctx context.Context, id int64) (entity.Product, error)
	FindAll(ctx context.Context) ([]entity.Product, error)
}

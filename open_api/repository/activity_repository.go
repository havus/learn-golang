package repository

import (
	"context"
	"database/sql"
	"open_api/model/domain"
)

type ActivityRepository interface {
	Save(ctx context.Context, 		tx *sql.Tx, activity domain.Activity) domain.Activity
	Update(ctx context.Context, 	tx *sql.Tx, activity domain.Activity) domain.Activity
	Delete(ctx context.Context, 	tx *sql.Tx, activity domain.Activity)
	FindById(ctx context.Context, tx *sql.Tx, activityId int) (domain.Activity, error)
	FindAll(ctx context.Context, 	tx *sql.Tx) []domain.Activity
}

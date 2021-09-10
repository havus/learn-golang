package repository

import (
	"context"
	"database/sql"
	"open_api/model/domain"
	"open_api/helper"
)

type ActivityRepositoryImpl struct {
}

func (repository *ActivityRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, activity domain.Activity) domain.Activity {
	query := "INSERT INTO activity(name, status) VALUES(?, ?)"
	result, err := tx.ExecContext(ctx, query, activity.Name, activity.Status)
	helper.PanicIfError(err)

	id, errId := result.LastInsertId()
	helper.PanicIfError(errId)

	activity.Id = int(id)
	return activity
}

func (repository *ActivityRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, activity domain.Activity) domain.Activity {
	query := "UPDATE activity SET name = ?, status = ? where id = ?"
	_, err := tx.ExecContext(ctx, query, activity.Name, activity.Status, activity.Id)
	helper.PanicIfError(err)

	return activity
}

func (repository *ActivityRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, activity domain.Activity) {
	query := "DELETE from activity where id = ?"
	_, err := tx.ExecContext(ctx, query, activity.Id)
	helper.PanicIfError(err)
}

func (repository *ActivityRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, activityId int) domain.Activity, error {
	query 		:= "SELECT id, name, status FROM activity WHERE id = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, query, activityId)
	helper.PanicIfError(err)

	activity := domain.Activity{}

	if rows.Next() {
		errRow := rows.Scan(&activity.Id, &activity.Name, &activity.Status)
		helper.PanicIfError(errRow)

		return activity, nil
	} else {
		return activity, errors.New("activity with id " + strconv.Itoa(int(id)) + " not found!")
	}
}

func (repository *ActivityRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Activity {
	query 		:= "SELECT id, name, status FROM activity"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)

	var activities []domain.Activity

	for rows.Next() {
		activity := entity.Activity{}
		errRow := rows.Scan(&activity.Id, &activity.Name, &activity.Status)
		helper.PanicIfError(errRow)

		activities = append(activities, activity) 
	}

	return activities
}
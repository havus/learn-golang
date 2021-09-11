package service

import (
	"context"
	"database/sql"
	"open_api/exception"
	"open_api/helper"
	"open_api/model/domain"
	"open_api/model/web"
	"open_api/repository"

	"github.com/go-playground/validator/v10"
)

type ActivityServiceImpl struct {
	ActivityRepository repository.ActivityRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewActivityService(ActivityRepository repository.ActivityRepository, DB *sql.DB, Validate *validator.Validate) ActivityService {
	return &ActivityServiceImpl{
		ActivityRepository: ActivityRepository,
		DB:                 DB,
		Validate:           Validate,
	}
}

func (service *ActivityServiceImpl) Create(ctx context.Context, request web.ActivityCreateRequest) web.ActivityResponse {
	errValidation := service.Validate.Struct(request)
	helper.PanicIfError(errValidation)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	activity := domain.Activity{
		// Id
		Name:   request.Name,
		Status: request.Status,
	}

	activity = service.ActivityRepository.Save(ctx, tx, activity)

	return helper.ToActivityResponse(activity)
}

func (service *ActivityServiceImpl) Update(ctx context.Context, request web.ActivityUpdateRequest) web.ActivityResponse {
	errValidation := service.Validate.Struct(request)
	helper.PanicIfError(errValidation)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	activity, errFind := service.ActivityRepository.FindById(ctx, tx, request.Id)
	if errFind != nil {
		panic(exception.NewNotFoundError(errFind.Error()))
	}

	activity.Name = request.Name
	activity.Status = request.Status

	activity = service.ActivityRepository.Update(ctx, tx, activity)

	return helper.ToActivityResponse(activity)
}

func (service *ActivityServiceImpl) Delete(ctx context.Context, activityId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	activity, errFind := service.ActivityRepository.FindById(ctx, tx, activityId)
	if errFind != nil {
		panic(exception.NewNotFoundError(errFind.Error()))
	}

	service.ActivityRepository.Delete(ctx, tx, activity)
}

func (service *ActivityServiceImpl) FindById(ctx context.Context, activityId int) web.ActivityResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	activity, errFind := service.ActivityRepository.FindById(ctx, tx, activityId)
	if errFind != nil {
		panic(exception.NewNotFoundError(errFind.Error()))
	}

	return helper.ToActivityResponse(activity)
}

func (service *ActivityServiceImpl) FindAll(ctx context.Context) []web.ActivityResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	activities := service.ActivityRepository.FindAll(ctx, tx)

	// var activityResponses []web.ActivityResponse

	// for _, activity := range activities {
	// 	activityResponses = append(activityResponses, helper.ToActivityResponse(activity))
	// }

	// return activityResponses
	return helper.ToActivityResponses(activities)
}

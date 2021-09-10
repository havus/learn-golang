package service

import (
	"context"
	"database/sql"
	"open_api/model/web"
	"open_api/repository"
	"open_api/helper"
	"open_api/model/domain"
)

type ActivityServiceImpl struct {
	ActivityRepository 	repository.ActivityRepository
	DB 									*sql.DB
	Validate						*validator.Validate
}

func (service *ActivityServiceImpl) Create(ctx context.Context, request web.ActivityCreateRequest) web.ActivityResponse {
	errValidation := service.Validate.Struct(request)
	helper.PanicIfError(errValidation)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	activity := domain.Activity{
		// Id
		Name: request.Name,
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

	activity, err := service.ActivityRepository.FindById(ctx, tx, request.Id)
	PanicIfError(err) // if not found

	activity.Name 	= request.Name
	activity.Status = request.Status

	activity = service.ActivityRepository.Update(ctx, tx, activity)

	return helper.ToActivityResponse(activity)
}

func (service *ActivityServiceImpl) Delete(ctx context.Context, activityId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	activity, err := service.ActivityRepository.FindById(ctx, tx, activityId)
	PanicIfError(err) // if not found

	service.ActivityRepository.Delete(ctx, tx, activityId)
}

func (service *ActivityServiceImpl) FindById(ctx context.Context, activityId int) web.ActivityResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	activity, err := service.ActivityRepository.FindById(ctx, tx, request.Id)
	PanicIfError(err) // if not found

	return helper.ToActivityResponse(activity)
}

func (service *ActivityServiceImpl) Create(ctx context.Context) []web.ActivityResponse {
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
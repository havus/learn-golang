package service

import (
	"context"
	"open_api/model/web"
)

type ActivityService interface {
	Create(ctx context.Context, request web.ActivityCreateRequest) web.ActivityResponse
	Update(ctx context.Context, request web.ActivityUpdateRequest) web.ActivityResponse
	Delete(ctx context.Context, activityId int)
	FindById(ctx context.Context, activityId int) web.ActivityResponse
	FindAll(ctx context.Context, activityId int) []web.ActivityResponse
}

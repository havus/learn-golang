package helper

import "open_api/model/domain"
import "open_api/model/web"

func ToActivityResponse(activity domain.Activity) web.ActivityResponse {
	return web.ActivityResponse{
		Id: activity.Id,
		Name: activity.Name,
		Status: activity.Status,
	}
}

func ToActivityResponses(activities []domain.Activity) []web.ActivityResponse {
	var activityResponses []web.ActivityResponse

	for _, activity := range activities {
		activityResponses = append(activityResponses, ToActivityResponse(activity))
	}

	return activityResponses
}

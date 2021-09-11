package controller

import (
	"strconv"
	"net/http"
	// "encoding/json"
	"github.com/julienschmidt/httprouter"
	"open_api/model/web"
	"open_api/service"
	"open_api/helper"
)

type ActivityControllerImpl struct {
	ActivityService service.ActivityService
}

func NewActivityController(activityService service.ActivityService) ActivityController {
	return &ActivityControllerImpl{
		ActivityService: activityService,
	}
}

func (controller *ActivityControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// decoder := json.NewDecoder(request.Body)

	// activityCreateRequest := web.ActivityCreateRequest{}
	// err := decoder.Decode(&activityCreateRequest)
	// helper.PanicIfError(err)
	// =============== refactor with ===============
	activityCreateRequest := web.ActivityCreateRequest{}
	helper.ReadFromRequestBody(request, &activityCreateRequest)

	activityResponse := controller.ActivityService.Create(request.Context(), activityCreateRequest)
	webResponse := web.WebResponse{
		Code: 	201,
		Status: "Ok",
		Data: 	activityResponse,
	}

	// writer.Header().Add("Content-Type", "application/json")

	// encoder := json.NewEncoder(writer)
	// errEncoder := encoder.Encode(webResponse)
	// helper.PanicIfError(errEncoder)
	// =============== refactor with ===============
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ActivityControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// decoder := json.NewDecoder(request.Body)

	// activityUpdateRequest := web.ActivityUpdateRequest{}
	// err := decoder.Decode(&activityUpdateRequest)
	// helper.PanicIfError(err)
	// =============== refactor with ===============
	activityUpdateRequest := web.ActivityUpdateRequest{}
	helper.ReadFromRequestBody(request, &activityUpdateRequest)

	paramId := params.ByName("activityId")
	intParamId, errConversion := strconv.Atoi(paramId) 
	helper.PanicIfError(errConversion)

	activityUpdateRequest.Id = intParamId

	activityResponse := controller.ActivityService.Update(request.Context(), activityUpdateRequest)
	webResponse := web.WebResponse{
		Code: 	200,
		Status: "Ok",
		Data: 	activityResponse,
	}

	// writer.Header().Add("Content-Type", "application/json")

	// encoder := json.NewEncoder(writer)
	// errEncoder := encoder.Encode(webResponse)
	// helper.PanicIfError(errEncoder)
	// =============== refactor with ===============
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ActivityControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paramId := params.ByName("activityId")
	intParamId, errConversion := strconv.Atoi(paramId) 
	helper.PanicIfError(errConversion)

	controller.ActivityService.Delete(request.Context(), intParamId)
	webResponse := web.WebResponse{
		Code: 	200,
		Status: "Ok",
	}

	// writer.Header().Add("Content-Type", "application/json")

	// encoder := json.NewEncoder(writer)
	// errEncoder := encoder.Encode(webResponse)
	// helper.PanicIfError(errEncoder)
	// =============== refactor with ===============
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ActivityControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paramId := params.ByName("activityId")
	intParamId, errConversion := strconv.Atoi(paramId) 
	helper.PanicIfError(errConversion)

	activityResponse := controller.ActivityService.FindById(request.Context(), intParamId)
	webResponse := web.WebResponse{
		Code: 	200,
		Status: "Ok",
		Data: 	activityResponse,
	}

	// writer.Header().Add("Content-Type", "application/json")

	// encoder := json.NewEncoder(writer)
	// errEncoder := encoder.Encode(webResponse)
	// helper.PanicIfError(errEncoder)
	// =============== refactor with ===============
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ActivityControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) { 
	activityResponses := controller.ActivityService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 	200,
		Status: "Ok",
		Data: 	activityResponses,
	}

	// writer.Header().Add("Content-Type", "application/json")

	// encoder := json.NewEncoder(writer)
	// errEncoder := encoder.Encode(webResponse)
	// helper.PanicIfError(errEncoder)
	// =============== refactor with ===============
	helper.WriteToResponseBody(writer, webResponse)
}

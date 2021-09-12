package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"open_api/app"
	"open_api/controller"
	"open_api/helper"
	"open_api/middleware"
	"open_api/model/domain"
	"open_api/repository"
	"open_api/service"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:rootroot@tcp(localhost:3306)/learn_golang_test?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	myValidator := validator.New()

	activityRepository 	:= repository.NewActivityRepository()
	activityService 		:= service.NewActivityService(activityRepository, db, myValidator)
	activityController 	:= controller.NewActivityController(activityService)

	router := app.NewRouter(activityController)

	return middleware.NewAuthMiddleware(router)
}

func truncateActivity(db *sql.DB) {
	db.Exec("TRUNCATE TABLE activity;")
}

// happy path
func TestCreateActivitySuccess(t *testing.T) {
	myDb 		:= setupTestDB()
	router 	:= setupRouter(myDb)

	defer truncateActivity(myDb)

	requestBody := strings.NewReader(`{"name": "mandi", "status": "todo"}`)
	request := httptest.NewRequest(http.MethodPost, "https://localhost:3000/api/activities", requestBody)
	request.Header.Add("x-api-key", "secret_key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)
	// fmt.Println(reflect.TypeOf(responseBody["code"]), responseBody["code"])

	assert := assert.New(t)

	assert.Equal(201, int(responseBody["code"].(float64)))
	assert.Equal("Ok", responseBody["status"])
	
	responseBodyData := responseBody["data"].(map[string]interface{})
	assert.Equal("mandi", responseBodyData["name"])
	assert.Equal("todo", responseBodyData["status"])
	// assert.Equal(t, 201, int(responseBody["code"].(float64)))
	// assert.Equal(t, "Ok", responseBody["status"])
	
	// responseBodyData := responseBody["data"].(map[string]interface{})
	// assert.Equal(t, "mandi", responseBodyData["name"])
	// assert.Equal(t, "todo", responseBodyData["status"])
}

func TestUpdateActivitySuccess(t *testing.T) {
	myDb 		:= setupTestDB()
	router 	:= setupRouter(myDb)

	defer truncateActivity(myDb)

	// Create activity ====== START ======
	tx, _ := myDb.Begin()
	activityRepository := repository.NewActivityRepository()
	activity := activityRepository.Save(context.Background(), tx, domain.Activity{Name: "Before Update", Status: "in progress"})
	tx.Commit()
	// Create activity ====== END ======

	requestBody := strings.NewReader(`{"name": "mandi", "status": "todo"}`)
	url	 				:= "https://localhost:3000/api/activities/" + strconv.Itoa(activity.Id)
	request 		:= httptest.NewRequest(http.MethodPut, url, requestBody)
	request.Header.Add("x-api-key", "secret_key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	fmt.Println(responseBody)

	assert := assert.New(t)
	
	assert.Equal(200, int(responseBody["code"].(float64)))
	assert.Equal("Ok", responseBody["status"])

	responseBodyData := responseBody["data"].(map[string]interface{})
	assert.Equal(activity.Id, int(responseBodyData["id"].(float64)))
	assert.Equal("mandi", responseBodyData["name"])
	assert.Equal("todo", responseBodyData["status"])
}

// sad path
func TestCreateActivityFailed(t *testing.T) {
	myDb 		:= setupTestDB()
	router 	:= setupRouter(myDb)

	defer truncateActivity(myDb)

	requestBody := strings.NewReader(`{"name": "", "status": "todos"}`)
	request := httptest.NewRequest(http.MethodPost, "https://localhost:3000/api/activities", requestBody)
	request.Header.Add("x-api-key", "secret_key")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])

	// responseBodyData := responseBody["data"].(map[string]interface{})
	assert.Contains(t, responseBody["data"], "Error:Field validation for 'Name'")
	// assert.Equal(t, "todo", responseBodyData["status"])
}
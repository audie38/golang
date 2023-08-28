package test

import (
	"database/sql"
	"encoding/json"
	"golang_api/config"
	"golang_api/controller"
	"golang_api/helper"
	"golang_api/middleware"
	"golang_api/repository"
	"golang_api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB{
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_db_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}

func setupRouter(db *sql.DB) http.Handler{
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := config.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func trucateCategory(db *sql.DB){
	db.Exec("TRUNCATE CATEGORY")
}

func TestCreateCategorySuccess(t *testing.T){
	db := setupTestDB()
	trucateCategory(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name":"Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, helper.CATEGORY_API_BASE_URL, requestBody)
	request.Header.Add(helper.CONTENT_TYPE, helper.APP_JSON)
	request.Header.Add(helper.API_KEY, helper.API_KEY_VAL)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)
	
	assert.Equal(t, 200, response.StatusCode, "HTTP Status Code Must be 200")
	assert.Equal(t, helper.RESPONSE_OK, responseBody["status"], "JSON Response Status Must be OK")
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"], "JSON Response Data Must be Gadget")
}

func TestCreateCategoryFailed(t *testing.T){
	db := setupTestDB()
	trucateCategory(db)
	router := setupRouter(db)
	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, helper.CATEGORY_API_BASE_URL, requestBody)
	request.Header.Add(helper.CONTENT_TYPE, helper.APP_JSON)
	request.Header.Add(helper.API_KEY, helper.API_KEY_VAL)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	var responseBody map[string]interface{}
	body, _ := io.ReadAll(response.Body)
	json.Unmarshal(body, &responseBody)
	
	assert.Equal(t, 400, response.StatusCode, "HTTP Status Code Must be 400")
	assert.Equal(t, helper.BAD_REQUEST_ERROR, responseBody["status"], "JSON Response Status Must be BAD REQUEST")
}

func TestUpdateCategorySuccess(t *testing.T){}

func TestUpdateCategoryFailed(t *testing.T){}

func TestGetCategorySuccess(t *testing.T){}

func TestGetCategoryFailed(t *testing.T){}

func TestGetListCategorySuccess(t *testing.T){}

func TestGetListCategoryFailed(t *testing.T){}

func TestDeleteCategorySuccess(t *testing.T){}

func TestDeleteCategoryFailed(t *testing.T){}

func TestUnauthorized(t *testing.T){}
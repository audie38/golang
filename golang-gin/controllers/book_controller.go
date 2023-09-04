package controllers

import (
	"golang-gin/middlewares"
	"golang-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

type BookRequest struct{
	Title string `json:"title"`
	Author string `json:"author"`
	Description string `json:"description"`
}

func (h handler) AddBook(ctx *gin.Context){
	body := BookRequest{}
	err := ctx.BindJSON(&body)
	if err != nil{
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book
	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	result := h.DB.Create(&book)
	if result.Error != nil{
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}

func (h handler) GetBook(ctx *gin.Context){
	var books []models.Book

	result := h.DB.Find(&books)
	if result.Error != nil{
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &books)
}

func (h handler) GetBookById(ctx *gin.Context){
	id := ctx.Param("id")
	var book models.Book

	result := h.DB.First(&book, id)
	if result.Error != nil{
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}

func (h handler) UpdateBook(ctx *gin.Context){
	id := ctx.Param("id")
	body := BookRequest{}

	err := ctx.BindJSON(&body)
	if err != nil{
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	result := h.DB.First(&book, id)
	if result.Error != nil{
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	h.DB.Save(&book)

	ctx.JSON(http.StatusOK, &book)
}

func (h handler) DeleteBook(ctx *gin.Context){
	id := ctx.Param("id")
	body := BookRequest{}

	err := ctx.BindJSON(&body)
	if err != nil{
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	result := h.DB.First(&book, id)
	if result.Error != nil{
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)
	ctx.Status(http.StatusOK)
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB){
	h := &handler{
		DB: db,
	}

	routes := router.Group("/api/book")
	routes.Use(middlewares.JwtAuthMiddleware())
	routes.POST("/", h.AddBook)
	routes.GET("/", h.GetBook)
	routes.GET("/:id", h.GetBookById)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
}
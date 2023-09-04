package controllers

import (
	"golang-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userHandler struct {
	DB *gorm.DB
}

type UserRequest struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func hashPassword(password string) (string, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		return "", err
	}
	return string(hashedPassword), nil
}

func verifyPassword(password string, hashedPassword string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (h userHandler)RegisterUser(c *gin.Context){
	body := UserRequest{}
	err := c.BindJSON(&body)
	if err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newUser := models.User{}
	newUser.Username = body.Username
	hashed, err := hashPassword(body.Password)
	if err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newUser.Password = hashed

	result := h.DB.Create(&newUser)
	if result.Error != nil{
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, &newUser)
}

func (h userHandler)Login(c *gin.Context){
	body := UserRequest{}
	err := c.BindJSON(&body)
	if err != nil{
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	existingUser := models.User{}
	result := h.DB.Model(models.User{}).Where("username = ?", body.Username).Take(&existingUser)
	if result.Error != nil{
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	err = verifyPassword(body.Password, existingUser.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword{
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
	})
}

func RegisterUserRoutes(router *gin.Engine, db *gorm.DB){
	h := &userHandler{
		DB: db,
	}

	routes := router.Group("/api/user")
	routes.POST("/regis", h.RegisterUser)
	routes.POST("/login", h.Login)
}
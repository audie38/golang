package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func GenerateToken(userId int64) (string, error) {
	JWT_SECRET := viper.Get("API_SECRET").(string)
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWT_SECRET))
}

func ExtractToken(c *gin.Context) string{
	token := c.Request.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2{
		return strings.Split(token, " ")[1]
	}
	return ""
}

func TokenValid(c *gin.Context) error{
	JWT_SECRET := viper.Get("API_SECRET").(string)
	tokenString := ExtractToken(c)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_SECRET), nil
	})

	if err != nil{
		return err
	}

	return nil
}
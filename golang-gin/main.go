package main

import (
	"golang-gin/config"
	"golang-gin/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("PG_URL").(string)

	router := gin.Default()
	dbHandler := config.Init(dbUrl)

	router.GET("/", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"port": port,
			"dbUrl": dbUrl,
		})
	})

	controllers.RegisterRoutes(router, dbHandler)

	router.Run(port)
}
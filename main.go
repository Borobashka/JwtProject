package main

import (
	"jwtProject/controllers"
	"jwtProject/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main(){
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
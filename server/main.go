package main

import (
	"github.com/dylankilkenny/watch-cash/server/controllers"
	"github.com/dylankilkenny/watch-cash/server/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	db.Init()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "hello world",
		})
	})
	router.POST("/signup", user.CreateUser)
	router.Run()
}

package main

import (
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dylankilkenny/watch-cash/server/controllers"
	"github.com/dylankilkenny/watch-cash/server/db"
	"github.com/gin-gonic/gin"
)

const secretkey = "verysecretkey1995"

func main() {
	router := gin.Default()
	db.Init()

	public := router.Group("/api")
	private := router.Group("/api/private")

	public.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  200,
			"message": "hello world",
		})
	})
	public.POST("/signup", user.SignUpUser)
	public.POST("/login", user.LoginUser)
	private.Use(Auth(secretkey))
	private.POST("/subscribe", user.SubscribeAddress)
	router.Run()
}

func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			c.AbortWithError(401, err)
		}
	}
}

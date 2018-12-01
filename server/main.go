package main

import (
	"net/http"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/dylankilkenny/watch-cash/bitsocket"
	"github.com/dylankilkenny/watch-cash/server/db"
	"github.com/dylankilkenny/watch-cash/server/user"
	"github.com/gin-gonic/gin"
)

const secretkey = "verysecretkey1995"

func main() {
	go bitsocket.Stream()

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
	private.Use(auth(secretkey))
	private.POST("/subscribe", user.SubscribeAddress)
	router.Run()

}

func auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			return b, nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "invalid token",
			})
		}
	}
}

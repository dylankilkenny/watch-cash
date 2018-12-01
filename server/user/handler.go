package user

import (
	"fmt"
	"log"
	"net/http"

	addressModel "github.com/dylankilkenny/watch-cash/server/address/model"
	"github.com/dylankilkenny/watch-cash/server/db"
	userModel "github.com/dylankilkenny/watch-cash/server/user/model"
	"github.com/dylankilkenny/watch-cash/server/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func SubscribeAddress(c *gin.Context) {
	db := db.GetDB()
	var address addressModel.Address

	if err := c.BindJSON(&address); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ID, err := jwt.Validate(c)
	fmt.Println("here")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "invalid token",
		})
		return
	}
	if err := db.Where("address = ? and user_id = ?", address.Address, ID).First(&address).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "already subscribed to address",
		})
		return
	}
	userId, err := uuid.FromString(ID)
	if err != nil {
		log.Fatal(err)
	}
	address.UserID = userId
	db.Create(&address)
	c.JSON(http.StatusOK, &address)
}

func LoginUser(c *gin.Context) {
	db := db.GetDB()
	var userLogin userModel.User
	var userDb userModel.User

	if err := c.BindJSON(&userLogin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Where("email = ?", userLogin.Email).First(&userDb).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "email does not exist",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(userLogin.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "password is incorrect",
		})
		return
	}

	token, err := jwt.Token(userDb.ID.String())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "invalid token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"token":     token,
		"firstname": userDb.FirstName,
	})
}

func SignUpUser(c *gin.Context) {
	var user userModel.User
	var db = db.GetDB()

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := db.Where("email = ?", user.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email already exists",
		})
		return
	}

	db.Create(&user)
	c.JSON(http.StatusOK, &user)
}

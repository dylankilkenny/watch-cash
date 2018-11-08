package user

import (
	"log"
	"net/http"

	"github.com/dylankilkenny/watch-cash/server/db"
	"github.com/dylankilkenny/watch-cash/server/models"
	"github.com/dylankilkenny/watch-cash/server/util/crypto"
	"github.com/dylankilkenny/watch-cash/server/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

func SubscribeAddress(c *gin.Context) {
	db := db.GetDB()
	var address models.Address

	if err := c.BindJSON(&address); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ID, err := jwt.Validate(c)
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
	var userLogin models.User
	var userDb models.User

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

	if err := crypto.Compare([]byte(userDb.Password), []byte(userLogin.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "password is incorrect",
		})
		return
	}

	token, err := jwt.Token(userDb.ID.String())
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"token":     token,
		"firstname": userDb.FirstName,
	})
}

func SignUpUser(c *gin.Context) {
	var user models.User
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

	password, err := crypto.Encrypt([]byte(user.Password))
	user.Password = string(password)
	if err != nil {
		log.Fatal(err)
	}

	db.Create(&user)
	c.JSON(http.StatusOK, &user)
}

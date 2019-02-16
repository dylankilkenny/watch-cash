package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dylankilkenny/watch-cash/server/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type subscribedAddressesResponse struct {
	CreatedAt string `json:"created_at"`
	Address   string `json:"address"`
}

type token struct {
	Token string `json:"token" binding:"required"`
}

func SubscribeToAddress(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		fmt.Println("Failed to fetch db")
	}
	var address Address

	if err := c.BindJSON(&address); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ID, err := jwt.Validate(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"code":   "invalid_token",
			"detail": "token is not valid",
		})
		return
	}

	if err := db.Where("address = ? and user_id = ?", address.Address, ID).First(&address).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"code":   "already_subscribed",
			"detail": "already subscribed to address",
		})
		return
	}

	userID, err := uuid.FromString(ID)
	if err != nil {
		log.Fatal(err)
	}
	address.UserID = userID
	db.Create(&address)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"detail": "address subscribed",
	})
}

func GetSubscribedAddresses(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		fmt.Println("Failed to fetch db")
	}
	var subscribedAddresses []Address
	ID, err := jwt.Validate(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"code":   "invalid_token",
			"detail": "token is not valid",
		})
		return
	}

	if err := db.Select("address, created_at").Where("user_id = ? and deleted_at is NULL", ID).Find(&subscribedAddresses).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	}

	var addresses []subscribedAddressesResponse
	for _, element := range subscribedAddresses {
		addresses = append(addresses, subscribedAddressesResponse{
			CreatedAt: element.CreatedAt.Format("01-02-2006"),
			Address:   element.Address,
		})

	}
	if len(addresses) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"detail": "No addresses found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"addresses": addresses,
	})
}

func RemoveSubscribedAddress(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		fmt.Println("Failed to fetch db")
	}
	var address Address

	if err := c.BindJSON(&address); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ID, err := jwt.Validate(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"code":   "invalid_token",
			"detail": "token is not valid",
		})
		return
	}

	if err := db.Joins("JOIN users ON addresses.user_id = ?", ID).Where("addresses.address = ?", address.Address).First(&address).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"detail": "address not found",
		})
		return
	}

	db.Delete(&address)

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"detail": "address removed",
	})
}

func LoginUser(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		fmt.Println("Failed to fetch db")
	}
	var userLogin User
	var userDb User

	if err := c.BindJSON(&userLogin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"detail": "email/password is missing",
		})
		return
	}

	if err := db.Where("email = ?", userLogin.Email).First(&userDb).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": 404,
			"detail": "email does not exist",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(userLogin.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"code":   "pass_wrong",
			"detail": "password is incorrect",
		})
		return
	}

	token, err := jwt.Token(userDb.ID.String())
	if err != nil {
		fmt.Printf("Error creating token for user %v", userDb.ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"token":     token,
		"firstname": userDb.FirstName,
	})
}

func ValidateToken(c *gin.Context) {
	var token token

	if err := c.BindJSON(&token); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"detail": "token is missing",
		})
		return
	}

	_, err := jwt.ValidateString(token.Token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"code":   "invalid_token",
			"detail": "token is not valid",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"code":   "valid_token",
			"detail": "token is valid",
		})
	}
}

func SignUpUser(c *gin.Context) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		fmt.Println("Failed to fetch db")
	}
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"code":   "invalid_form",
			"detail": "invalid form",
		})
		return
	}

	if err := db.Where("email = ?", user.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"code":   "email_taken",
			"error":  "email already exists",
		})
		return
	}

	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"detail": "user sign up success",
	})
}

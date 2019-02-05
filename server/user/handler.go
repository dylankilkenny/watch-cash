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

func SubscribeToAddress(c *gin.Context) {
	db := db.GetDB()
	var address addressModel.Address

	if err := c.BindJSON(&address); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ID, err := jwt.Validate(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": 401,
			"code":   "invalid_token",
			"detail": "token is not valid",
		})
		return
	}

	if !db.Joins("JOIN users ON addresses.user_id = ?", ID).Where("addresses.address = ?", address.Address).First(&address).RecordNotFound() {
		if address.DeletedAt != nil {
			address.DeletedAt = nil
			db.Save(&address)
			c.JSON(http.StatusOK, &address)
			return
		}
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

type subscribedAddressesResponse struct {
	CreatedAt string `json:"created_at"`
	Address   string `json:"address"`
}

func GetSubscribedAddresses(c *gin.Context) {
	db := db.GetDB()
	var subscribedAddresses []addressModel.Address
	ID, err := jwt.Validate(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
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
		fmt.Println(element)
		addresses = append(addresses, subscribedAddressesResponse{
			CreatedAt: element.CreatedAt.Format("01-02-2006"),
			Address:   element.Address,
		})

	}
	if len(addresses) == 0 {
		c.JSON(http.StatusOK, gin.H{
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
	db := db.GetDB()
	var address addressModel.Address

	if err := c.BindJSON(&address); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ID, err := jwt.Validate(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
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
	db := db.GetDB()
	var userLogin userModel.User
	var userDb userModel.User

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
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"code":   "expired_token",
			"detail": "token is expired",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    200,
		"token":     token,
		"firstname": userDb.FirstName,
	})
}

type token struct {
	Token string `json:"token" binding:"required"`
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
	fmt.Println(token)

	_, err := jwt.ValidateString(token.Token)
	if err != nil {
		fmt.Println(err)
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
	var user userModel.User
	var db = db.GetDB()

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

func SubscribedUsers(address string) ([]userModel.User, error) {
	var users []userModel.User
	var db = db.GetDB()

	if err := db.Joins("JOIN addresses ON addresses.user_id = users.id").Where("addresses.address = ?", address).Find(&users).Error; err == nil {
		return users, err
	}
	return users, nil
}

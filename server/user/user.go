package user

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dylankilkenny/watch-cash/server/util/jwt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
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
			"status": 401,
			"errors": gin.H{
				"title":  "Internal server error",
				"detail": err.Error(),
			},
			"data": make([]string, 0),
		})
		return
	}

	ID, err := jwt.Validate(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"errors": gin.H{
				"title":  "Invalid Token",
				"detail": "Token provided in auth header is not valid",
			},
			"data": make([]string, 0),
		})
		return
	}

	if err := db.Where("address = ? and user_id = ?", address.Address, ID).First(&address).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": gin.H{
				"title":  "Already Subscribed",
				"detail": "Already subscribed to the provided address",
			},
			"data": make([]string, 0),
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
		"errors": make([]string, 0),
		"data":   make([]string, 0),
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
			"errors": gin.H{
				"title":  "Invalid Token",
				"detail": "Token provided in auth header is not valid",
			},
			"data": make([]string, 0),
		})
		return
	}

	if err := db.Select("address, created_at").Where("user_id = ? and deleted_at is NULL", ID).Find(&subscribedAddresses).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": 401,
			"errors": gin.H{
				"title":  "Internal server error",
				"detail": err.Error(),
			},
			"data": make([]string, 0),
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
			"errors": gin.H{
				"title":  "No addresses found",
				"detail": "The user is not subscribed to any addresses",
			},
			"data": make([]string, 0),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data": gin.H{
			"addresses": addresses,
		},
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
			"status": 401,
			"errors": gin.H{
				"title":  "Internal server error",
				"detail": err.Error(),
			},
			"data": make([]string, 0),
		})
		return
	}

	ID, err := jwt.Validate(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"errors": gin.H{
				"title":  "Invalid Token",
				"detail": "Token provided in auth header is not valid",
			},
			"data": make([]string, 0),
		})
		return
	}

	if err := db.Joins("JOIN users ON addresses.user_id = ?", ID).Where("addresses.address = ?", address.Address).First(&address).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": gin.H{
				"title":  "Address not found",
				"detail": "No matching address was found",
			},
			"data": make([]string, 0),
		})
		return
	}

	db.Delete(&address)

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   make([]string, 0),
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
			"errors": gin.H{
				"title":  "Email or password missing",
				"detail": "Email or password missing",
			},
			"data": make([]string, 0),
		})
		return
	}

	if err := db.Where("email = ?", userLogin.Email).First(&userDb).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": 404,
			"errors": gin.H{
				"title":  "Email Not Found",
				"detail": "No matching email address was found",
			},
			"data": make([]string, 0),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(userLogin.Password)); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"errors": gin.H{
				"title":  "Wrong Password",
				"detail": "password is incorrect",
			},
			"data": make([]string, 0),
		})
		return
	}

	token, err := jwt.Token(userDb.ID.String())
	if err != nil {
		fmt.Printf("Error creating token for user %v", userDb.ID)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data": gin.H{
			"token":     token,
			"firstname": userDb.FirstName,
		},
	})
}

func ValidateToken(c *gin.Context) {
	var token token

	if err := c.BindJSON(&token); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": gin.H{
				"title":  "Token Missing",
				"detail": "No token was supplied",
			},
			"data": make([]string, 0),
		})
		return
	}

	_, err := jwt.ValidateString(token.Token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status": 401,
			"errors": gin.H{
				"title":  "Invalid Token",
				"detail": "token is not valid",
			},
			"data": make([]string, 0),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"errors": make([]string, 0),
			"data":   make([]string, 0),
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
			"errors": gin.H{
				"title":  "Invalid Form",
				"detail": "The submitted form is not valid",
			},
			"data": make([]string, 0),
		})
		return
	}

	if err := db.Where("email = ?", user.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": gin.H{
				"title":  "Email Taken",
				"detail": "email already exists",
			},
			"data": make([]string, 0),
		})
		return
	}

	db.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"errors": make([]string, 0),
		"data":   make([]string, 0),
	})
}

func SubscribedUsers(c *gin.Context) {

	apiKey := c.Request.Header.Get("API_KEY")

	if apiKey == os.Getenv("WATCHCASHAPIKEY") {
		db, ok := c.MustGet("db").(*gorm.DB)
		if !ok {
			fmt.Println("Failed to fetch db")
		}
		var users []User
		var address Address

		if err := c.BindJSON(&address); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if db.Joins("JOIN addresses ON addresses.user_id = users.id").Where("addresses.address = ?", address.Address).Find(&User{}).RecordNotFound() {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": 400,
				"errors": gin.H{
					"title":  "No users",
					"detail": "no subscribed users were found",
				},
				"data": make([]string, 0),
			})
			return
		}

		if err := db.Joins("JOIN addresses ON addresses.user_id = users.id").Where("addresses.address = ?", address.Address).Find(&users).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": 500,
				"errors": gin.H{
					"title":  "Something went wrong",
					"detail": err.Error(),
				},
				"data": make([]string, 0),
			})
			return
		}

		response := make([]map[string]string, 0)
		for _, u := range users {
			response = append(response, map[string]string{
				"name":  u.FirstName,
				"email": u.Email,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"errors": make([]string, 0),
			"data":   response,
		})

		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"status": 200,
		"errors": gin.H{
			"title":  "API key",
			"detail": "api key is not valid",
		},
		"data": make([]string, 0),
	})
}

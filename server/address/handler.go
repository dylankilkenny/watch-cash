package address

import (
	"net/http"

	"github.com/dylankilkenny/watch-cash/server/db"
	"github.com/gin-gonic/gin"
)

func Addresses() {
	if err := db.Where("address = ? and user_id = ?", address.Address, ID).First(&address).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "already subscribed to address",
		})
		return
	}
}

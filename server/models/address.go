package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Address struct {
	BaseModel
	Address string    `json:"address"`
	UserID  uuid.UUID `json:"user_id"`
}

func (address *Address) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Println("Uuid err")
		panic(err)
	}
	scope.SetColumn("ID", uuid.String())
	return nil
}

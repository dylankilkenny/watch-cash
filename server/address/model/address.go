package model

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}

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

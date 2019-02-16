package user

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"deletedAt,omitempty"`
}

type User struct {
	BaseModel
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Addresses []Address `gorm:"foreignkey:UserID";json:"addresses"`
}

type Address struct {
	BaseModel
	Address string    `json:"address"`
	UserID  uuid.UUID `json:"user_id"`
}

func (user *User) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Println("Uuid err")
		panic(err)
	}
	scope.SetColumn("ID", uuid.String())
	if pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err == nil {
		scope.SetColumn("Password", string(pw))
	}

	return nil
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

func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", "check")
	return nil
}

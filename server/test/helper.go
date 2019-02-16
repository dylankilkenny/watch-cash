package test

import (
	"fmt"
	"log"

	"github.com/dylankilkenny/watch-cash/server/user"
	"github.com/dylankilkenny/watch-cash/server/util/jwt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //db
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func InitDB() *gorm.DB {

	userDb := "dylankilkenny"
	password := ""
	host := "localhost"
	port := "5432"
	dbname := "test"

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		userDb,
		password,
		host,
		port,
		dbname,
	)

	db, err = gorm.Open("postgres", dbinfo)

	if err != nil {
		log.Println("Failed to connect to testing database")
		panic(err)
	}
	log.Println("Testing Database connected")

	CreateCleanDB()

	return db
}

func CreateUser() *user.User {
	user := user.User{FirstName: "Dylan", LastName: "Kilkenny", Email: "email@email.com", Password: "password"}
	db.Create(&user)
	return &user
}

func CreateCleanDB() {
	db.DropTableIfExists(&user.User{}, &user.Address{})

	if !db.HasTable(&user.User{}) {
		db.CreateTable(&user.User{})
	}

	if !db.HasTable(&user.Address{}) {
		db.CreateTable(&user.Address{})
	}
}

func GetToken(u *user.User) (token string) {
	token, err := jwt.Token(u.ID.String())
	if err != nil {
		fmt.Printf("Error creating token for user %v", u.ID)
	}
	return
}

func AddAddress(u *user.User) {
	var address user.Address
	address.UserID = u.ID
	address.Address = "qrwd7tucj2l6rjcgv5cr2n4t8ws83ghsjqpar98qpt"
	db.Create(&address)
}

package db

import (
	"fmt"
	"log"
	"os"

	addressModel "github.com/dylankilkenny/watch-cash/server/address/model"
	userModel "github.com/dylankilkenny/watch-cash/server/user/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var err error

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Init creates a connection to mysql database and
// migrates any new models
func Init(debug bool) {
	user := getEnv("PG_USER", "dylankilkenny")
	password := getEnv("PG_PASSWORD", "")
	host := getEnv("PG_HOST", "localhost")
	port := getEnv("PG_PORT", "5432")
	database := getEnv("PG_DB", "watchcash")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db, err = gorm.Open("postgres", dbinfo)
	db.LogMode(debug)

	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")
	if !db.HasTable(&userModel.User{}) {
		err := db.CreateTable(&userModel.User{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	if !db.HasTable(&addressModel.Address{}) {
		err := db.CreateTable(&addressModel.Address{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&addressModel.Address{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

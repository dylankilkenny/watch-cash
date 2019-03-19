package db

import (
	"fmt"
	"log"
	"os"

	"github.com/dylankilkenny/watch-cash/server/user"
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

// Init creates a connection to postgres database and
// migrates any new models
func Init(debug bool) {
	userDb := getEnv("PG_USER", "")
	password := getEnv("PG_PASSWORD", "")
	host := getEnv("PG_HOST", "")
	port := getEnv("PG_PORT", "5432")
	database := getEnv("PG_DB", "")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		userDb,
		password,
		host,
		port,
		database,
	)

	fmt.Println(dbinfo)

	db, err = gorm.Open("postgres", dbinfo)
	db.LogMode(debug)

	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}
	log.Println("Database connected")
	if !db.HasTable(&user.User{}) {
		err := db.CreateTable(&user.User{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	if !db.HasTable(&user.Address{}) {
		err := db.CreateTable(&user.Address{})
		if err != nil {
			log.Println("Table already exists")
		}
	}

	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&user.Address{})
}

//GetDB ...
func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	db.Close()
}

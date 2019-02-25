package main

import (
	"log"
	"os"

	"github.com/dylankilkenny/watch-cash/server/db"
	"github.com/dylankilkenny/watch-cash/server/router"
	"github.com/joho/godotenv"
)

func main() {
	debugMode := false
	logging := true

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.Init(debugMode)
	db := db.GetDB()

	r := router.SetupRouter(db, logging)
	r.Run(":" + os.Getenv("PORT"))
}

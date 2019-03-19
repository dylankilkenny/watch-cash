package main

import (
	"os"

	"github.com/dylankilkenny/watch-cash/server/db"
	"github.com/dylankilkenny/watch-cash/server/router"
)

func main() {
	debugMode := false
	logging := true

	// if value, ok := os.LookupEnv(key); !ok {
	// 	log.Fatal("main() -> Error loading .env file")
	// }
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Println(err)
	// 	log.Fatal("main() -> Error loading .env file")
	// }

	db.Init(debugMode)
	db := db.GetDB()

	r := router.SetupRouter(db, logging)
	r.Run(":" + os.Getenv("PORT"))
}

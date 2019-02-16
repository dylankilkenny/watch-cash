package main

import (
	"github.com/dylankilkenny/watch-cash/server/db"
	"github.com/dylankilkenny/watch-cash/server/router"
)

func main() {
	debugMode := false
	logging := true
	db.Init(debugMode)
	db := db.GetDB()
	r := router.SetupRouter(db, logging)
	r.Run(":3001")

}

package main

import (
	"github.com/ilhamnyto/twinyto/config"
	"github.com/ilhamnyto/twinyto/pkg/database"
)

func main() {
	config.LoadConfig(".env")

	dbsql, err := database.ConnectDB()

	if err != nil {
		panic(err)
	}

	db := database.NewDatabase()
	db = db.SetSQL(dbsql)
	
}
package main

import (
	"fmt"

	"github.com/ilhamnyto/twinyto/apps"
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
	
	port := fmt.Sprintf("0.0.0.0:%s",config.GetString(config.PORT))

	factory := apps.NewRouterFactory(db)
	router, err := factory.Create(apps.Router_GIN, port)

	if err != nil {
		panic(err)
	}

	executor := apps.NewRouterExecutor()
	executor.Execute(router)
}
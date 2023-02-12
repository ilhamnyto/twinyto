package database

import (
	"database/sql"
	"fmt"

	"github.com/ilhamnyto/twinyto/config"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	var (
		db_host = config.GetString(config.DB_HOST)
		db_port = config.GetString(config.DB_PORT)
		db_user = config.GetString(config.DB_USER)
		db_password = config.GetString(config.DB_PASSWORD)
		db_name = config.GetString(config.DB_NAME)
	)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db_host, db_port, db_user, db_password, db_name,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
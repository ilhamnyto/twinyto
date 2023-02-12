package database

import "database/sql"

type Database struct {
	DbSQL *sql.DB
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) SetSQL(db *sql.DB) *Database {
	d.DbSQL = db
	return d
}


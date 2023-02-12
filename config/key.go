package config

import "os"

const (
	DB_HOST = "DB_HOST"
	DB_PORT = "DB_PORT"
	DB_USER = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_NAME = "DB_NAME"
	PORT = "PORT"
	SECRET_KEY = "SECRET_KEY"
)

func GetString(key string) string {
	return os.Getenv(key)
}
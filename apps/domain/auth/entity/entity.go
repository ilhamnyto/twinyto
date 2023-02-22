package entity

import "time"

type User struct {
	Id        int
	Username  string
	Email     string
	ImgUrl    string
	Password  string
	CreatedAt time.Time
}
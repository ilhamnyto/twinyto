package entity

import "time"

type Post struct {
	Id        int
	UserId    int
	Body      string
	CreatedAt time.Time
}

type UserPost struct {
	Id int
	Username string
	Body string
	CreatedAt time.Time
}
package entity

import "time"

type Follow struct {
	UserId      int
	FollowingId int
	CreatedAt   time.Time
}


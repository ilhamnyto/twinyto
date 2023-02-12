package token

import "time"

type Payload struct {
	UserId int
	Expired time.Time
}
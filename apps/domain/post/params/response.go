package params

import (
	"time"

	"github.com/ilhamnyto/twinyto/apps/domain/post/entity"
)

type UserPostResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Body      string `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *UserPostResponse) ParseFromModel(m *entity.UserPost) {
	u.Id = m.Id
	u.Username = m.Username
	u.Body = m.Body
	u.CreatedAt = m.CreatedAt
}
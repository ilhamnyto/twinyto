package params

import (
	"time"

	"github.com/ilhamnyto/twinyto/apps/domain/post/entity"
)

type UserCreatePostRequest struct {
	UserId    int `json:"user_id"`
	Body      string `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

type UserDeletePostRequest struct {
	UserId    int `json:"user_id"`
	PostId int `json:"post_id"`
	CreatedAt string `json:"created_at"`
}

func (u *UserCreatePostRequest) ParseToModel() *entity.Post {
	return &entity.Post{
		UserId: u.UserId,
		Body: u.Body,
		CreatedAt: u.CreatedAt,
	}
}

func (u *UserDeletePostRequest) ParseToModel() *entity.Post {
	return &entity.Post{
		Id: u.PostId,
		UserId: u.PostId,
	}
}
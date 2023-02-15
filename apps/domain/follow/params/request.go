package params

import (
	"time"

	"github.com/ilhamnyto/twinyto/apps/domain/follow/entity"
)

type UserFollowRequest struct {
	UserId      int `json:"user_id"`
	FollowingId int `json:"following_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserUnfollowRequest struct {
	UserId int `json:"user_id"`
	FollowingId int `json:"following_id"`
}

func (u *UserFollowRequest) ParseToModel() *entity.Follow {
	return &entity.Follow{
		UserId: u.UserId,
		FollowingId: u.FollowingId,
		CreatedAt: u.CreatedAt,
	}
}
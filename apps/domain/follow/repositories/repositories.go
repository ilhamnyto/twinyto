package repositories

import (
	"context"

	"github.com/ilhamnyto/twinyto/apps/domain/follow/entity"
)

type FollowRepo interface {
	Create(ctx context.Context, req *entity.Follow) error
	Delete(ctx context.Context, userId int, followingId int) error
}
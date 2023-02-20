package repositories

import (
	"context"

	"github.com/ilhamnyto/twinyto/apps/domain/profile/entity"
)

type ProfileRepo interface {
	FindById(ctx context.Context, userId int) (*entity.Profile, error)
	FindByUsername(ctx context.Context, username string) (*entity.Profile, error)
	FindUser(ctx context.Context, username string) ([]*entity.Profile, error)
	GetAllUser(ctx context.Context) ([]*entity.Profile, error)
	GetFollower(ctx context.Context, userId int) ([]*entity.Profile, error)
	GetUserIdByUsername(ctx context.Context, username string) (int, error)
}
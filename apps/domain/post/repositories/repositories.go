package repositories

import (
	"context"
	"time"

	"github.com/ilhamnyto/twinyto/apps/domain/post/entity"
)

type PostRepo interface {
	Create(ctx context.Context, post *entity.Post) error
	Delete(ctx context.Context, post *entity.Post) error
	GetPost(ctx context.Context, username string, postId int, createdAt time.Time) (*entity.UserPost, error)
	GetUserPosts(ctx context.Context, username string) ([]*entity.UserPost, error)
	GetAllPost(ctx context.Context) ([]*entity.UserPost, error)
}
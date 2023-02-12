package repositories

import (
	"context"

	"github.com/ilhamnyto/twinyto/apps/domain/auth/entity"
)

type AuthRepo interface {
	Create(ctx context.Context, user *entity.User) error
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	UpdatePassword(ctx context.Context, password string, userid int) error
	GetPassword(ctx context.Context, userid int) (string, error)
}
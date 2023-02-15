package usecase

import (
	"github.com/ilhamnyto/twinyto/apps/domain/follow/repositories"
	"github.com/ilhamnyto/twinyto/apps/domain/follow/services"
)

type followSvc struct {
	repo repositories.FollowRepo
}

func NewFollowSvc(repo repositories.FollowRepo) services.FollowSvc {
	return &followSvc{
		repo: repo,
	}
}

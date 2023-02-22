package usecase

import (
	"github.com/ilhamnyto/twinyto/apps/domain/post/repositories"
	"github.com/ilhamnyto/twinyto/apps/domain/post/services"
)

type postSvc struct {
	repo repositories.PostRepo
}

func NewPostSvc(repo repositories.PostRepo) services.PostSvc {
	return &postSvc {
		repo: repo,
	}
}
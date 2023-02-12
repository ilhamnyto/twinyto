package usecase

import (
	"github.com/ilhamnyto/twinyto/apps/domain/auth/repositories"
	"github.com/ilhamnyto/twinyto/apps/domain/auth/services"
)

type authSvc struct {
	repo repositories.AuthRepo
}

func NewAuthSvc(repo repositories.AuthRepo) services.AuthSvc {
	return &authSvc {
		repo: repo,
	}
}




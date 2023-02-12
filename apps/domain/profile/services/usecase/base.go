package usecase

import (
	"github.com/ilhamnyto/twinyto/apps/domain/profile/repositories"
	"github.com/ilhamnyto/twinyto/apps/domain/profile/services"
)

type profileSvc struct {
	repo repositories.ProfileRepo
}

func NewProfileSvc(repo repositories.ProfileRepo) services.ProfileSvc {
	return &profileSvc{
		repo: repo,
	}
}


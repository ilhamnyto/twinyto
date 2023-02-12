package usecase

import (
	"context"

	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/profile/params"
)

func (s *profileSvc) MyProfile(ctx context.Context, userId int) (*params.UserProfileResponse, *response.CustomError) {
	
}

func (s *profileSvc) UserProfile(ctx context.Context, username string) (*params.UserProfileResponse, *response.CustomError) {

}

func (s *profileSvc) SearchProfile(ctx context.Context, req *params.UserProfileRequest) ([]*params.UserProfileResponse, *response.CustomError) {

}

func (s *profileSvc) UserProfileList(ctx context.Context) ([]*params.UserProfileResponse, *response.CustomError) {

}
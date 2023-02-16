package services

import (
	"context"

	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/profile/params"
)

type ProfileSvc interface {
	MyProfile(ctx context.Context, userId int) (*params.UserProfileResponse, *response.CustomError)
	UserProfile(ctx context.Context, username string) (*params.UserProfileResponse, *response.CustomError)
	SearchProfile(ctx context.Context, req *params.UserProfileRequest) ([]*params.UserProfileResponse, *response.CustomError)
	UserProfileList(ctx context.Context) ([]*params.UserProfileResponse, *response.CustomError)
	UserFollowerList(ctx context.Context, userId int)([]*params.UserProfileResponse, *response.CustomError)
}
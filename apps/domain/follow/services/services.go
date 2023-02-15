package services

import (
	"context"

	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/follow/params"
)

type FollowSvc interface {
	Follow(ctx context.Context, req *params.UserFollowRequest) *response.CustomError
	Unfollow(ctx context.Context, req *params.UserUnfollowRequest) *response.CustomError
}
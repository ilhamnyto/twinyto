package usecase

import (
	"context"
	"time"

	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/follow/params"
)

func (f *followSvc) Follow(ctx context.Context, req *params.UserFollowRequest) *response.CustomError {
	user := req.ParseToModel()
	user.CreatedAt = time.Now()
	
	err := f.repo.Create(ctx, user)

	if err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
	
}

func (f *followSvc) Unfollow(ctx context.Context, req *params.UserUnfollowRequest) *response.CustomError {
	err := f.repo.Delete(ctx, req.UserId, req.FollowingId)

	if err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}
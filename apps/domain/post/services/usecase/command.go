package usecase

import (
	"context"

	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/post/params"
)

func (p *postSvc) CreatePost(ctx context.Context, req *params.UserCreatePostRequest) *response.CustomError {
	post := req.ParseToModel()

	if err := p.repo.Create(ctx, post); err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}

func (p *postSvc) DeletePost(ctx context.Context, req *params.UserDeletePostRequest) *response.CustomError {
	post := req.ParseToModel()
	
	if err := p.repo.Delete(ctx, post); err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}
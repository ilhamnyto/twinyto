package usecase

import (
	"context"
	"time"

	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/post/params"
)

func (p *postSvc) GetPost(ctx context.Context, username string, postId int, createdAt time.Time) (*params.UserPostResponse, *response.CustomError) {
	return nil, nil
}

func (p *postSvc) GetUserPost(ctx context.Context, username string) ([]*params.UserPostResponse, *response.CustomError) {
	return nil, nil
}

func (p *postSvc) GetAllPost(ctx context.Context) ([]*params.UserPostResponse, *response.CustomError) {
	return nil, nil
}
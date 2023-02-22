package services

import (
	"context"
	"time"

	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/post/params"
)

type PostSvc interface {
	CreatePost(ctx context.Context, req *params.UserCreatePostRequest) *response.CustomError
	DeletePost(ctx context.Context, req *params.UserDeletePostRequest) *response.CustomError
	GetPost(ctx context.Context, username string, postId int, createdAt time.Time) (*params.UserPostResponse, *response.CustomError)
	GetUserPost(ctx context.Context, username string) ([]*params.UserPostResponse, *response.CustomError)
	GetAllPost(ctx context.Context) ([]*params.UserPostResponse, *response.CustomError)
}
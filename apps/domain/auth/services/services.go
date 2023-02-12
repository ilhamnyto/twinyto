package services

import (
	"context"

	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/auth/params"
)

type AuthSvc interface {
	Login(ctx context.Context, req *params.UserLoginRequest) (*params.UserLoginResponse, *response.CustomError)
	Register(ctx context.Context, req *params.UserRegisterRequest) (*response.CustomError)
	ResetPassword(ctx context.Context, req *params.UserPasswordResetRequest, userId int) (*response.CustomError)
}
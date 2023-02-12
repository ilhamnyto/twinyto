package usecase

import (
	"context"
	"database/sql"

	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/auth/params"
	"github.com/ilhamnyto/twinyto/pkg/encryption"
	"github.com/ilhamnyto/twinyto/pkg/token"
)

func (s *authSvc) Login(ctx context.Context, req *params.UserLoginRequest) (*params.UserLoginResponse, *response.CustomError) {
	user, err := s.repo.FindByUsername(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NotFoundError()
		}

		return nil, response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	err = encryption.ValidatePassword(user.Password, req.Password)

	if err != nil {
		return nil, response.GeneralErrorWithAdditionalInfo("Wrong Password.")
	}

	tokenString, err := token.GenerateToken(user.Id)

	if err != nil {
		return nil, response.GeneralErrorWithAdditionalInfo(err.Error())
	}

	return &params.UserLoginResponse{
		AccessToken: tokenString,
	}, nil
}
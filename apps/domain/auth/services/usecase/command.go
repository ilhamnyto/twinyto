package usecase

import (
	"context"
	"database/sql"

	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/auth/params"
)

func (s *authSvc) Register(ctx context.Context, req *params.UserRegisterRequest) *response.CustomError {
	if req.Password != req.ConfirmPassword {
		return response.GeneralErrorWithAdditionalInfo("Password did'nt match.")
	}
	user := req.ParseToModel()
	user.ImgUrl = "https://fastly.picsum.photos/id/630/500/500.jpg?hmac=_e8WfDqIZfqQ0doa8XEoc4JEw2SQq2ud7QplFmfS6Ag"
	err := s.repo.Create(ctx, user)

	if err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}

func (s *authSvc) ResetPassword(ctx context.Context, req *params.UserPasswordResetRequest, userId int) *response.CustomError {
	if req.Password != req.ConfirmPassword {
		return response.GeneralErrorWithAdditionalInfo("Password didn't match.")
	}

	currPass, err := s.repo.GetPassword(ctx, userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return response.NotFoundError()
		}

		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	if currPass == req.Password {
		return response.GeneralErrorWithAdditionalInfo("Password can't be the same as before.")
	}

	err = s.repo.UpdatePassword(ctx, req.Password, userId)

	if err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}
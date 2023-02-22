package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/auth/params"
	"github.com/ilhamnyto/twinyto/pkg/encryption"
)

func (s *authSvc) Register(ctx context.Context, req *params.UserRegisterRequest) *response.CustomError {
	if req.Password != req.ConfirmPassword {
		return response.GeneralErrorWithAdditionalInfo("Password did'nt match.")
	}

	if len (req.Password) < 5 {
		return response.GeneralErrorWithAdditionalInfo("password should have at least 5 characters.")
	}

	loc, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		return response.GeneralErrorWithAdditionalInfo(err.Error())
	}

	createdat := time.Now().In(loc)
	req.CreatedAt = createdat

	user := req.ParseToModel()
	user.ImgUrl = "https://fastly.picsum.photos/id/630/500/500.jpg?hmac=_e8WfDqIZfqQ0doa8XEoc4JEw2SQq2ud7QplFmfS6Ag"

	hash, err := encryption.HashPassword(req.Password)

	if err != nil {
		return response.GeneralErrorWithAdditionalInfo(err.Error())
	}

	user.Password = hash

	err = s.repo.Create(ctx, user)

	if err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}

func (s *authSvc) ResetPassword(ctx context.Context, req *params.UserPasswordResetRequest, userId int) *response.CustomError {
	if req.Password != req.ConfirmPassword {
		return response.GeneralErrorWithAdditionalInfo("Password didn't match.")
	}

	if len (req.Password) < 5 {
		return response.GeneralErrorWithAdditionalInfo("password should have at least 5 characters.")
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

	hash, err := encryption.HashPassword(req.Password)

	if err != nil {
		return response.GeneralErrorWithAdditionalInfo(err.Error())
	}
	req.Password = hash
	err = s.repo.UpdatePassword(ctx, req.Password, userId)

	if err != nil {
		return response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	return nil
}
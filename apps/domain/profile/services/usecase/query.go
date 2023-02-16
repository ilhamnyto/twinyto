package usecase

import (
	"context"
	"database/sql"

	"github.com/ilhamnyto/twinyto/apps/commons/response"
	"github.com/ilhamnyto/twinyto/apps/domain/profile/params"
)

func (s *profileSvc) MyProfile(ctx context.Context, userId int) (*params.UserProfileResponse, *response.CustomError) {
	result, err := s.repo.FindById(ctx, userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NotFoundError()
		}

		return nil, response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	user := params.UserProfileResponse{}

	user.ParseFromModel(result)

	return &user, nil

}

func (s *profileSvc) UserProfile(ctx context.Context, username string) (*params.UserProfileResponse, *response.CustomError) {
	result, err := s.repo.FindByUsername(ctx, username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NotFoundError()
		}

		return nil, response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	user := params.UserProfileResponse{}
	user.ParseFromModel(result)

	return &user, nil
}

func (s *profileSvc) SearchProfile(ctx context.Context, req *params.UserProfileRequest) ([]*params.UserProfileResponse, *response.CustomError) {
	result, err := s.repo.FindUser(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NotFoundError()
		}

		return nil, response.NotFoundErrorWithAdditionalInfo(err.Error())
	}

	var users []*params.UserProfileResponse

	for _, user := range result {
		tempUser := new(params.UserProfileResponse)
		tempUser.ParseFromModel(user)

		users = append(users, tempUser)
	}

	if users == nil {
		users = []*params.UserProfileResponse{}
	}

	return users, nil
}

func (s *profileSvc) UserProfileList(ctx context.Context) ([]*params.UserProfileResponse, *response.CustomError) {
	result, err := s.repo.GetAllUser(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NotFoundError()
		}

		return nil, response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	var users []*params.UserProfileResponse

	for _, user := range result {
		tempUser := new(params.UserProfileResponse)
		tempUser.ParseFromModel(user)

		users = append(users, tempUser)
	}

	if users == nil {
		users = []*params.UserProfileResponse{}
	}

	return users, nil
}

func (s *profileSvc) UserFollowerList(ctx context.Context, userId int) ([]*params.UserProfileResponse, *response.CustomError) {
	result, err := s.repo.GetFollower(ctx, userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, response.NotFoundError()
		}

		return nil, response.RepositoryErrorWithAdditionalInfo(err.Error())
	}

	var users []*params.UserProfileResponse

	for _, user := range result {
		tempUser := new(params.UserProfileResponse)
		tempUser.ParseFromModel(user)

		users = append(users, tempUser)
	}

	if users == nil {
		users = []*params.UserProfileResponse{}
	}

	return users, nil
}
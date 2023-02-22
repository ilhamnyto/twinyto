package params

import (
	"time"

	"github.com/ilhamnyto/twinyto/apps/domain/auth/entity"
)

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRegisterRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	ImgUrl string `json:"image_url"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *UserRegisterRequest) ParseToModel() *entity.User {
	return &entity.User{
		Username: u.Username,
		Email: u.Email,
		Password: u.Password,
		ImgUrl: u.ImgUrl,
		CreatedAt: u.CreatedAt,
	}
}

type UserPasswordResetRequest struct {
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}


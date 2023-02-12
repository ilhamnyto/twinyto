package params

import "github.com/ilhamnyto/twinyto/apps/domain/profile/entity"

type UserProfileResponse struct {
	Username string `json:"username"`
	Email string `json:"email"`
	ImgUrl string `json:"img_url"`
}

func (u *UserProfileResponse) ParseFromModel(user *entity.Profile) {
	u.Email = user.Email
	u.Username = user.Username
	u.ImgUrl = user.ImgUrl
}

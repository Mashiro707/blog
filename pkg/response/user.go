package response

import "blog/internal/model"

type UserInfoWithToken struct {
	model.User
	Token string `json:"token"`
}

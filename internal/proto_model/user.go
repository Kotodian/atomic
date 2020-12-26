package proto_model

import (
	"atomic/atomic_model"
	"atomic/atomic_server/proto/user"
)

func User(user *user.User) *atomic_model.User {
	return &atomic_model.User{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Phone:    user.Phone,
	}
}

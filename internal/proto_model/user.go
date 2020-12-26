package proto_model

import (
	user2 "atomic/atomic_model/user"
	"atomic/atomic_server/proto/user"
)

func User(user *user.User) *user2.User {
	return &user2.User{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Phone:    user.Phone,
	}
}

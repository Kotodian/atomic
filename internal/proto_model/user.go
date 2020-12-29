package proto_model

import (
	model "atomic/atomic_model/user"
	proto "atomic/atomic_proto/user"
)

func User(user *proto.User) *model.User {
	return &model.User{
		Id:       user.Id,
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
		Phone:    user.Phone,
	}
}

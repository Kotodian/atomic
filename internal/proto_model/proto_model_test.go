package proto_model

import (
	"atomic/atomic_model/user"
	proto "atomic/atomic_proto/user"
	"testing"
)

func TestProtoToModel(t *testing.T) {
	var u = &user.User{}
	err := ProtoToModel(&proto.User{
		Id:       1,
		Username: "lqk",
		Nickname: "lqk",
		Email:    "lqk",
		Phone:    "lqk",
		Password: "lqk",
	}, u)
	if err != nil {
		t.Error(err)
	}
	t.Log(u)
}

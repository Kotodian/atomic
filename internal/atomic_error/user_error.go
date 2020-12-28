package atomic_error

import "errors"

var (
	UserNotExists = errors.New("用户不存在")
	PasswordWrong = errors.New("密码错误")
)

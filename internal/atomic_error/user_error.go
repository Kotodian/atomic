package atomic_error

import "errors"

var (
	PasswordWrong = errors.New("密码错误")
)

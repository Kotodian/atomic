package user_err

import "errors"

var (
	PasswordWrong = errors.New("密码错误")
)

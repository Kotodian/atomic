package atomic_error

import "errors"

var (
	ErrUserNotExists = errors.New("用户不存在")
	ErrPasswordWrong = errors.New("密码错误")
	ErrUserCreate    = errors.New("用户创建失败")
	ErrUserLogin     = errors.New("用户登录失败")
	ErrUserNotOnline = errors.New("该用户当前不是登录状态")
	ErrUserLogout    = errors.New("用户退出失败")
)

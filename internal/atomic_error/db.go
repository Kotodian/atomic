package atomic_error

import "errors"

var (
	ErrConnect = errors.New("数据库连接错误")
	ErrUpdate  = errors.New("数据更新失败")
)

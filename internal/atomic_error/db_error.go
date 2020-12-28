package atomic_error

import "errors"

var (
	ErrConnect = errors.New("数据库连接错误")
)

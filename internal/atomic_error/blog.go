package atomic_error

import "errors"

var (
	ErrCreateBlog = errors.New("博客创建失败")
)

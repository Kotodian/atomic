package atomic_error

import "errors"

var (
	ErrCreateBlog = errors.New("创建博客失败")
	ErrDeleteBlog = errors.New("删除博客失败")
)

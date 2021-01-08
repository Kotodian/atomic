package atomic_error

import "errors"

var (
	ErrCreateCategory = errors.New("创建分类失败")
	ErrUpdateCategory = errors.New("更新分类失败")
	ErrDeleteCategory = errors.New("删除分类失败")
)

package atomic_error

import "errors"

var (
	ErrCollectionCreate = errors.New("收藏失败")
	ErrCollectionDelete = errors.New("取消收藏失败")
)

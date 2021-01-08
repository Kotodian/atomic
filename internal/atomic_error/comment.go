package atomic_error

import "errors"

var (
	ErrCommentCreate = errors.New("创建评论失败")
	ErrCommentDelete = errors.New("删除评论失败")
)

package atomic_model

import (
	"context"
	"gorm.io/gorm"
)

type Blog interface {
	// 创建者 返回用户
	Proposer(context.Context, *gorm.DB) User
	// 创建时间 返回时间戳
	ProposeTime(context.Context, *gorm.DB) int64
	// 是否拥有节点
	HasNode(context.Context, *gorm.DB) bool
	// 创建节点
	CreateNode(context.Context, *gorm.DB, Node) error
	// 插入 int64是用户id
	Insert(context.Context, *gorm.DB, int64) error
	// 更新
	Update(context.Context, *gorm.DB) error
	// 删除
	Delete(context.Context, *gorm.DB) error
}

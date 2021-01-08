package atomic_model

import (
	"context"
	"gorm.io/gorm"
)

type Blog interface {
	GetId() int64
	// 创建时间 返回时间戳
	ProposeTime(context.Context, *gorm.DB) int64
	// 是否拥有节点
	HasNode(context.Context, *gorm.DB) bool
	// 创建节点
	CreateNode(context.Context, *gorm.DB, Node) error
	// 插入 string是用户名
	Insert(context.Context, *gorm.DB, string) error
	// 更新
	Update(context.Context, *gorm.DB) error
	// 删除
	Delete(context.Context, *gorm.DB) error
	// 增加或减少收藏数
	UpdateCollection(context.Context, *gorm.DB, bool) error
	// 增加或减少点赞数
	UpdateKudos(context.Context, *gorm.DB, bool) error
	// 增加浏览数
	UpdateBrowse(context.Context, *gorm.DB) error
}

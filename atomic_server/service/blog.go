package service

import "context"

type Blog interface {
	// 创建者 返回用户
	Proposer(context.Context) User
	// 创建时间 返回时间戳
	ProposeTime(context.Context) int64
	// 是否拥有节点
	HasNode(context.Context) bool
	// 创建节点
	CreateNode(context.Context, Node) error
}

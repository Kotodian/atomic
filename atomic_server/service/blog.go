package service

type Blog interface {
	// 创建者 返回用户
	Proposer() User
	// 创建时间 返回时间戳
	ProposeTime() int64
	// 是否拥有节点
	HasNode() bool
	// 创建节点
	CreateNode(Node) error
}

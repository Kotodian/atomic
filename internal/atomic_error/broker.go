package atomic_error

import "errors"

var (
	ErrBrokerConnect = errors.New("消息组件无法连接")
)

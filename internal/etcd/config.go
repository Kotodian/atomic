package etcd

import "time"

var Etcds = []string{
	"192.168.2.235:12379",
	"192.168.2.235:22379",
	"192.168.2.235:32379",
}

const (
	DefaultTimeout = 5 * time.Second
)

package atomic_store

type dataSource int

const (
	MYSQL dataSource = iota
	PGSQL
	Mongodb
)

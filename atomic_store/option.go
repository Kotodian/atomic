package atomic_store

type Option interface {
	DefaultOption()
}

type MysqlOption struct {
	// dsn
	dsn string
	// 默认string长度
	DefaultStringSize uint
	// 禁用datetime精度
	DisableDatetimePrecision bool
}

func (m *MysqlOption) DefaultOption() {
	m.DisableDatetimePrecision = mysqlDisableDatetimePrecision
	m.DefaultStringSize = mysqlDefaultStringSize
	m.dsn = mysqlDefaultDSN
}

const (
	mysqlDefaultDSN               = "root:7k164173520@tcp(127.0.0.1:3306)/atomic?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDefaultStringSize        = 255
	mysqlDisableDatetimePrecision = true
)

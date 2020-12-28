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
	mysqlDefaultDSN               = ""
	mysqlDefaultStringSize        = 255
	mysqlDisableDatetimePrecision = true
)

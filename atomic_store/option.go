package atomic_store

type Option struct {
}

type MysqlOption struct {
	Option
	// 默认string长度
	DefaultStringSize uint
	// 禁用datetime精度
	DisableDatetimePrecision bool
}

const (
	defaultStringSize        = 255
	disableDatetimePrecision = true
)

func DefaultMysqlOption() *MysqlOption {
	return &MysqlOption{
		DefaultStringSize:        defaultStringSize,
		DisableDatetimePrecision: disableDatetimePrecision,
	}
}

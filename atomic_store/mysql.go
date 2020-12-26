package atomic_store

import (
	"atomic/internal/log"
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	option *MysqlOption
}

func (m *Mysql) Init(ctx context.Context, dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DefaultStringSize:        m.option.DefaultStringSize,
		DisableDatetimePrecision: m.option.DisableDatetimePrecision,
	}), &gorm.Config{})
	if err != nil {
		log.Error(err.Error(), ctx)
		return nil, err
	}
	return db, nil
}

// 参数设置
func (m *Mysql) Option(opt *MysqlOption) {
	m.option = opt
}

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

func (m *Mysql) DatabaseDefaultOption() {
	m.option.DefaultOption()
}

func (m *Mysql) Init(ctx context.Context) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      m.option.dsn,
		DefaultStringSize:        m.option.DefaultStringSize,
		DisableDatetimePrecision: m.option.DisableDatetimePrecision,
	}), &gorm.Config{})
	if err != nil {
		log.Error(err, ctx)
		return nil, err
	}
	return db, nil
}

package atomic_store

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Mysql struct {
	option *MysqlOption
}

func (m *Mysql) DatabaseDefaultOption() {
	m.option = new(MysqlOption)
	m.option.DefaultOption()
}

func (m *Mysql) Init(ctx context.Context) (*gorm.DB, error) {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		Colorful:      true,
		LogLevel:      logger.Info,
	})
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      m.option.dsn,
		DefaultStringSize:        m.option.DefaultStringSize,
		DisableDatetimePrecision: m.option.DisableDatetimePrecision,
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

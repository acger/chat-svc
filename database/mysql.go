package database

import (
	"github.com/acger/chat-svc/internal/config"
	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewMysql(c *config.Config) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               c.Datasource,
		DefaultStringSize: 255,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		logx.Error("mysql connect fail")
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	migrate(db)

	return db
}

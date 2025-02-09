package database

import (
	"context"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/kohge2/upsdct-server/config"
)

type DB struct {
	conn *gorm.DB
}

func GetDB(dsn string) (*DB, error) {
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	// sqliteでも良さそうだがデータ型指定したかったりするのでmysqlを使う(使い慣れてるというのもある)
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, err
	}
	return &DB{conn: db}, nil
}

func (db DB) GetNewTxnOrContext(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(config.ContextKeyTxn).(*gorm.DB)
	if ok {
		return tx
	}
	return db.conn
}

func NewDB(gormDB *gorm.DB) *DB {
	return &DB{conn: gormDB}
}

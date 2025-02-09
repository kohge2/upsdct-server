package database

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/kohge2/upsdct-server/config"
)

type Transaction struct {
	db *gorm.DB
}

func NewTransaction(dsn string) (*Transaction, error) {
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		return nil, err
	}
	return &Transaction{db: db}, nil
}

func (t *Transaction) RunTxn(ctx context.Context, txFunc func(ctx context.Context) error) (err error) {
	txn := t.db.Begin()
	if txn.Error != nil {
		return txn.Error
	}

	ctx = t.set(ctx, txn)

	defer func() {
		if p := recover(); p != nil {
			txn.Rollback()
			stackTrace := debug.Stack()
			err = fmt.Errorf("panic caught: %v\nstack trace: %s", p, stackTrace)
			log.Printf("Panic during transaction: %v\nStack trace: %s", p, stackTrace)
			return
		} else if err != nil {
			txn.Rollback()
			return
		} else {
			txn.Commit()
			return
		}
	}()

	err = txFunc(ctx)
	return err
}

func (t *Transaction) set(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, config.ContextKeyTxn, tx)
}

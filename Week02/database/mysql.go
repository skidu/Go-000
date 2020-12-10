package database

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db                *gorm.DB
	ErrRecordNotFound = errors.New(`record not found`)
)

func InitMySQL() {
	var (
		err error
		dsn = `... fake mysql conn dsn`
	)
	db, err = gorm.Open(`mysql`, dsn)
	if err != nil {
		panic(`connect to mysql failed: ` + err.Error())
	}
}

func GetConn() *gorm.DB {
	return db
}

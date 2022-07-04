package mysql

import (
	"blog/config"
	"blog/pkg/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Database interface {
	Open(conn string) (db *gorm.DB, err error)
}

type Mysql struct {
}

func (m *Mysql) Setup() {
	var (
		err error
		db  Database
	)
	db = new(Mysql)
	MysqlConn := config.MysqlConfig
	DB, err = db.Open(MysqlConn.DSN)
	if err != nil {
		zap.ErrorLog(err)
		panic("database connect error")
	} else {
		zap.InfoLog("database connect success")
	}
}

func (m *Mysql) Open(conn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(conn))
	if err != nil {
		zap.ErrorLog(err)
		return nil, err
	}
	return db, err
}

package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

// Init 初始化
func Init(dsn string) (err error) {
	DB, err = sqlx.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 查看是否连接成功
	err = DB.Ping()
	if err != nil {
		return err
	}
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

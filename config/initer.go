package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(mysql MysqlConf) (db *sqlx.DB) {
	connStr := mysql.Username + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/" + mysql.DBName + "?" + mysql.Config
	db, err := sqlx.Open("mysql", connStr)
	if err != nil {
		fmt.Printf("mysql connect error: %s", err.Error())
		return
	}
	if err = db.Ping(); err != nil {
		fmt.Printf("mysql connect error: %s", err.Error())
		return
	}
	return db
}

func InitGormMysql(mysqlconf MysqlConf) (db *gorm.DB) {
	connStr := mysqlconf.Username + ":" + mysqlconf.Password + "@tcp(" + mysqlconf.Host + ":" + mysqlconf.Port + ")/" + mysqlconf.DBName + "?" + mysqlconf.Config

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		fmt.Printf("gorm conn fail")
		return
	}
	return
}

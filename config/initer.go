package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func InitMysql(mysql Mysql) (db *sqlx.DB) {
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

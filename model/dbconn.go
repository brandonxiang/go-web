package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


func Init() (db *sqlx.DB) {
	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/example")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	fmt.Println(database)
	return database
}
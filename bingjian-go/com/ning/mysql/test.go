package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	dsName := ""
	db,err := sql.Open("mysql",dsName)
	if err != nil {
		fmt.Println(err)
		return
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(3)
	db.SetConnMaxLifetime(7 * time.Hour)
}

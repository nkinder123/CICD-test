package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// CICD流水线搭建
func main() {
	//链接数据库u
	db, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/kratos")
	if err != nil {
		fmt.Printf("sql Open has error:%v", err)
		log.Fatal("sql open has error:", err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Printf("sql Open has error:%v", err)
		log.Fatal("sql ping has error ")
		return
	}

	var version string
	row := db.QueryRow("SELECT VERSION()", &version)
	if err != nil {
		log.Fatal("db query has error")
		return
	}
	fmt.Printf("result:%v", row)
}

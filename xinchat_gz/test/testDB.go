package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func main() {

	DB, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/gim_gozero?charset=utf8mb4&parseTime=True&loc=Local")
	DB.SetMaxOpenConns(2000)
	DB.SetMaxIdleConns(1000)
	DB.SetConnMaxLifetime(time.Minute * 60) // mysql default conn timeout=8h, should < mysql_timeout
	err = DB.Ping()
	if err != nil {
		log.Fatalf("database init failed, err: ", err)
	}
	log.Println("mysql conn pool has initiated.")

}

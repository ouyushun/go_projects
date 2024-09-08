package main

import (
	"database/sql"
	"fmt"
	"github.com/edwingeng/wuid/mysql/wuid"
)

func main() {

	openDB := func() (*sql.DB, bool, error) {
		var db *sql.DB
		// ...
		dsn := fmt.Sprintf("root:12345678@tcp(127.0.0.1:3306)/gim_gozero?charset=utf8mb4&parseTime=True&loc=Local")
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, false, err
		}
		return db, true, nil
	}

	// Setup
	w := wuid.NewWUID("aaa", nil)
	err := w.LoadH28FromMysql(openDB, "wuid")
	if err != nil {
		panic(err)
	}

	// Generate
	for i := 0; i < 3; i++ {
		fmt.Printf("%#016x\n", w.Next())
	}
}

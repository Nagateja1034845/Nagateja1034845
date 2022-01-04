package main

import (
	"database/sql"
	"fmt"

	_ "githum.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("GO MySQL")

	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

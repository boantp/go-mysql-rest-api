package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:pasaribu@tcp(database.dev:3306)/shopee")
	checkErr(err)

	fmt.Println("You connected to your database.")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

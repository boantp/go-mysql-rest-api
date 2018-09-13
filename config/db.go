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

	// rows, err := DB.Query("SELECT tax_code_id, name FROM tax_code")
	// checkErr(err)

	// for rows.Next() {
	// 	var tax_code_id int
	// 	var name string
	// 	err = rows.Scan(&tax_code_id, &name)
	// 	checkErr(err)
	// 	fmt.Println(tax_code_id)
	// 	fmt.Println(name)
	// }
	fmt.Println("You connected to your database.")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

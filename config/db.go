package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "root:"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_IP_ADDR")+":"+os.Getenv("MYSQL_PORT")+")/shopee")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

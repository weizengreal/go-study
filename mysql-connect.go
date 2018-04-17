package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)



func main() {
	db, err := sql.Open("mysql","root:@127.0.0.1:3306/phalcon_test")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
}



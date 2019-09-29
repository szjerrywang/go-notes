package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Hello, world!")

	db, err := sql.Open("mysql",
		"didasoft:Jzhc201909@tcp(127.0.0.1:3306)/appbooks")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

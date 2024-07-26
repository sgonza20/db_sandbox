package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "mydatabase:password(127.0.0.1:3306)/mydatabase") // Open connection to MySQL DB

	if err != nil { // Check if an error occured
		panic(err) // If so, return the error
	}
	ping := db.Ping() // ping the database
	if ping != nil {  // check if ping returned an error
		fmt.Println("Error pinging database", ping) // If so, print error message
	}

}

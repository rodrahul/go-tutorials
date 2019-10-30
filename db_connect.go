package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "tempHughes:pi@s9m@tcp(10.19.129.63:3306)/lems")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	results, err := db.Query(`SELECT 
	Organization_Name
FROM
	Lms_Organization
WHERE 
	Organization_Status = 'Active'`)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var name string
		err = results.Scan(&name)
		if err != nil {
			fmt.Println("err after scanning", err)
		}
		fmt.Println(name)
	}

}

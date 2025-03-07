package model

import (
	"fmt"

	_ "github.com/jmoiron/sqlx"
)

// INIT DATABASE HERE
type Database struct {
	//use sqlx.DB if you want to connect database
	// db *sqlx.DB
	db string
}

func InitDatabase(driver, host, user, password, database string, port int) *Database {
	// setup database
	fmt.Println("=====Setting up database=====\n")
	fmt.Printf("driver: %s\n", driver)
	fmt.Printf("host: %s\n", host)
	fmt.Printf("user: %s\n", user)
	fmt.Printf("password: %s\n", password)
	fmt.Printf("database: %s\n", database)
	fmt.Printf("port: %d\n", port)
	fmt.Println("======Database Connected=====\n")
	
	return &Database{
		db: "Postgres",
	}

}

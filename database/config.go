package database

import "fmt"

var (
	dbUsername = "postgres"
	dbPassword = "postgres"
	dbHost     = "localhost"
	dbPort     = "5432"
	dbTable    = "postgres"
	pgConnStr  = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbTable, dbPassword)
)

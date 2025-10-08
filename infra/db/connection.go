package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(host, port, user, password, dbname string) string {
	return "user=postgres password=postgres dbname=ecommerce host=localhost port=5432 "
}

func NewConnection() (*sqlx.DB, error) {
	dbSource := GetConnectionString("localhost", "5432", "postgres", "postgres", "ecommerce")
	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}
	return db, nil
}

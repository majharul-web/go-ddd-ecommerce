package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(conf *config.DBConfig) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		conf.User, conf.Password, conf.DBName, conf.Host, conf.Port, conf.SSLMode)
}

func NewConnection(conf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(conf)
	db, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}
	return db, nil
}

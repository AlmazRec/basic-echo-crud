package database

import (
	"awesomeProject3/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func InitDB(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database)

	var err error
	db, err = sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return db
}

func GetDB() *sql.DB {
	return db
}

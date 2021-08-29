package database_learn

import (
	"fmt"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {
	fmt.Println("Connected to database sql")

	db, err := sql.Open("mysql", "root:rootroot@tcp(localhost:3306)/golang?parseTime=true")

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
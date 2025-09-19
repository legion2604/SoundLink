package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	// Подключение к БД
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/soundLink")
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Проверка соединения
	if err = DB.Ping(); err != nil {
		log.Fatal("Database ping failed: ", err)
	}
}

/*
id
name
surname
email
password

*/

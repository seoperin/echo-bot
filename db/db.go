package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// ConnectDB establishes a connection to the database and returns the DB instance
func ConnectDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
		return nil, err
	}

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
		return nil, err
	}

	fmt.Println("Успешно подключено к базе данных!")
	return db, nil
}

package repository

import (
	"database/sql"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
)

func Conn() *sql.DB {
	// Строка подключения
	connString := "server=localhost;port=1433;user id=sa;password=YourStrong!Passw0rd;database=master;encrypt=disable"

	// Открываем соединение
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Ошибка подключения:", err)
	}

	return db
}

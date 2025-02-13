package users

import (
	"back/cmd/hadlers/structs"
	"back/internal/repository"
	"database/sql"
	"log"
)

func CreatUser(req structs.RegisterRequest) error {
	db := repository.Conn()
	defer db.Close()
	query := "INSERT INTO USERS (first_name, last_name, email, hash_pass) VALUES (@first_name, @last_name, @email, @hash_pass);"
	log.Print(req)
	_, err := db.Exec(query,
		sql.Named("first_name", req.FirstName),
		sql.Named("last_name", req.LastName),
		sql.Named("email", req.Email),
		sql.Named("hash_pass", req.Password),
	)
	if err != nil {
		log.Fatal("Ошибка при вставке данных:", err)
	}

	return nil
}

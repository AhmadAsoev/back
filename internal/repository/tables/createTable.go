package tables

import (
	"back/cmd/hadlers/structs"
	"back/internal/repository"
	"database/sql"
	"log"
)

func CreateTable(req structs.TableRequest) error {
	db := repository.Conn()
	defer db.Close()
	query := "INSERT INTO ACTS (act_number, sender_name, sender_position, sender_organization, date, receiver_name, receiver_position,receiver_organization) VALUES (@act_number, @sender_name, @sender_position, @sender_organization, @date, @receiver_name, @receiver_position, @receiver_organization);"
	log.Print(req)
	_, err := db.Exec(query,
		sql.Named("act_number", req.ActNumber),
		sql.Named("sender_name", req.SenderName),
		sql.Named("sender_position", req.SenderPositon),
		sql.Named("sender_organization", req.SenderOrganization),
		sql.Named("date", req.Date),
		sql.Named("receiver_name", req.ReceiverName),
		sql.Named("receiver_position", req.ReceiverPosition),
		sql.Named("receiver_organization", req.ReceiverOrganization),
	)
	if err != nil {
		log.Fatal("Ошибка при вставке данных:", err)
	}

	return nil
}

package auth

import (
	"back/cmd/hadlers/structs"
	"back/internal/repository/tables"
	"log"
)

func CreateTable(req structs.TableRequest) error {

	if err := tables.CreateTable(req); err != nil {
		log.Print(err)
		return err
	}

	return nil
}

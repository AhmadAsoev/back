package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"

	"back/cmd/hadlers/structs"
	"back/internal/repository/users"
)

func CreatUser(req structs.RegisterRequest) (err error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print(err)
		return err
	}

	req.Password = string(hash)

	if err = users.CreatUser(req); err != nil {
		log.Print(err)
		return err
	}

	return nil
}

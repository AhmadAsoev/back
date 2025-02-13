package authorization

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"back/cmd/hadlers/structs"
)

func Login(c fiber.Ctx) error {
	var data structs.LoginRequest

	data.Email = c.Params("username")
	data.Password = c.Params("password")

	log.Print(data)

	return c.JSON("")
}

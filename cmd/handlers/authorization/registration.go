package autthorization

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v3"

	"back/cmd/handlers/structs"
)

func Registration(c fiber.Ctx) error {

	var data structs.RegistrationRequest

	if err := json.Unmarshal(c.Body(), &data); err != nil {
		log.Println("Ошибка при парсинге JSON:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Невозможно обработать JSON",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Данные успешно получены и записаны в лог",
		"data":    data,
	})
}

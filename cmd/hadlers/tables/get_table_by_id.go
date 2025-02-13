package tables

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v3"

	"back/cmd/hadlers/structs"
)

func GetTableByID(c fiber.Ctx) error {
	var data structs.RegisterRequest

	if err := json.Unmarshal(c.Body(), &data); err != nil {
		log.Println("Ошибка при парсинге JSON:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Невозможно обработать JSON",
		})
	}

	return c.JSON("")
}

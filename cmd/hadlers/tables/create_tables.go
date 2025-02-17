package tables

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v3"

	"back/cmd/hadlers/structs"
	"back/internal/auth"
)

func CreateTable(c fiber.Ctx) error {
	var data structs.TableRequest

	// body := c.Body()

	// log.Print(string(body))

	if err := json.Unmarshal(c.Body(), &data); err != nil {
		log.Println("Ошибка при парсинге JSON:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Невозможно обработать JSON",
		})
	}

	if err := auth.CreateTable(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Невозможно создать таблицу",
		})

	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Акт успешно добавлен!",
	})
}

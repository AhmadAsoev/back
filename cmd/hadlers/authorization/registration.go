package authorization

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v3"

	"back/cmd/hadlers/structs"
	"back/internal/auth"
)

func Registration(c fiber.Ctx) error {
	var data structs.RegisterRequest

	if err := json.Unmarshal(c.Body(), &data); err != nil {
		log.Println("Ошибка при парсинге JSON:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Невозможно обработать JSON",
		})
	}

	if err := auth.CreatUser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Невозможно создать пользователя",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Пользователь зарегестрировался!",
	})
}

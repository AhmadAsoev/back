package main

import (
	"log"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

type RequestData struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, 
		AllowMethods:     []string{"GET,POST,PUT,DELETE,OPTIONS"},
		AllowHeaders:     []string{"Content-Type,Authorization"},
		AllowCredentials: true,
	}))



	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/print", func(c fiber.Ctx) error {
		log.Print("Salom Ahmad")
		return nil
	})


	app.Post("/register", func(c fiber.Ctx) error {
		
		var data RequestData

		log.Println("Raw body: ", string(c.Body()) )

		if err := c.BodyParser(&data); err != nil {
			log.Println("Ошибка при парсинге JSON:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Невозможно обработать JSON",
			})
		}

		log.Println("Получены данные: FirstName=%s, LastName=%s\n,  Email=%s\n, Password= %s\n, ConfirmPassword=%s\n", data.FirstName, data.LastName, data.Email, data.Password, data.ConfirmPassword)

        return c.JSON(fiber.Map{
			"message": "Данные успешно получены и записаны в лог",
			"data": data,
		})
    })

	log.Fatal(app.Listen(":3000"))
}

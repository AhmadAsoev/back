package main

import (
	"log"

	"back/cmd/handlers/authorization"
	"back/cmd/handlers/health"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET,POST,PUT,DELETE,OPTIONS"},
		AllowHeaders:     []string{"Content-Type,Authorization"},
		AllowCredentials: true,
	}))

	//health
	app.Get("/health", health.Health)

	//registration
	app.Post("/register", autthorization.Registration)

	log.Fatal(app.Listen(":3000"))
}

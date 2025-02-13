package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"back/cmd/hadlers/authorization"
	"back/cmd/hadlers/health"
	"back/cmd/hadlers/tables"
)

func main() {
	
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET,POST,PUT,DELETE,OPTIONS"},
		AllowHeaders:     []string{"Content-Type,Authorization"},
		AllowCredentials: true,
	}))
	// Heath
	app.Get("/health", health.Health)

	// Auth
	app.Post("/register", authorization.Registration)
	app.Get("/login/:username/:password", authorization.Login)

	// Table
	app.Post("/tables", tables.CreateTable)
	app.Get("/tables/list", tables.GatTables)
	app.Get("/tables/:id", tables.GetTableByID)

	

	log.Fatal(app.Listen(":3000"))
}

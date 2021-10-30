package main

import (
	"github.com/gofiber/fiber/v2"
	"go-ambdassador/src/database"
	"go-ambdassador/src/routes"
)

func main() {
	database.Connect()
	database.AutoMigrate()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
}

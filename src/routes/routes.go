package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-ambdassador/src/controllers"
)

func Setup(app *fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")

	admin.Get("/register", controllers.Register)

}

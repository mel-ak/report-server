package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mel-ak/report-server/handlers"
)

func SetupUserRoutes(app *fiber.App) {
	app.Post("/users/register", handlers.RegisterUser)
	app.Post("/users/login", handlers.LoginUser)
}

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mel-ak/report-server/config"
	"github.com/mel-ak/report-server/database"
	"github.com/mel-ak/report-server/routes"
)

func main() {
	config.LoadEnv()

	// Custom config
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Report App v1.0.1",
	})

	app.Use(helmet.New())
	// Initialize default config
	app.Use(logger.New())
	// Get route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to the Report Backend!",
		})
	})

	database.ConnectDB()

	routes.SetupRoutes(app)

	// Listening to port 3000
	log.Fatal(app.Listen(":3000"))
}

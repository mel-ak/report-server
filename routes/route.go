// routes/routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mel-ak/report-server/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")        // /api
	v1 := api.Group("/v1")          // /api/v1
	reports := v1.Group("/reports") // /v1/reports

	reports.Get("/", handlers.GetReports)
	reports.Post("/", handlers.CreateReport)
	reports.Get("/:id", handlers.GetReport)
	reports.Put("/:id", handlers.UpdateReport)
	reports.Delete("/:id", handlers.DeleteReport)
}

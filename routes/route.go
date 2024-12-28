// routes/routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mel-ak/report-server/handlers"
	"github.com/mel-ak/report-server/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")        // /api
	v1 := api.Group("/v1")          // /api/v1
	reports := v1.Group("/reports") // /v1/reports

	reports.Get("/", middleware.AuthenticateJWT, handlers.GetReports)
	reports.Post("/", middleware.AuthenticateJWT, handlers.CreateReport)
	reports.Get("/:id", middleware.AuthenticateJWT, handlers.GetReport)
	reports.Put("/:id", middleware.AuthenticateJWT, handlers.UpdateReport)
	reports.Delete("/:id", middleware.AuthenticateJWT, handlers.DeleteReport)
}

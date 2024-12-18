package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/mel-ak/report-server/models"
)


func main(){	
	// Initialize the new fiber app
	app := fiber.New();
    
	// Get route
	app.Get("/", func(c fiber.Ctx) error{
		return c.JSON(fiber.Map{
			"message": "Welcome to the Report Backend!",
		})
	})

	api := app.Group("/api") // /api
	v1 := api.Group("/v1")        // /api/v1
	reports:= v1.Group("/reports") // /v1/reports

	// Create reports route
	reports.Post("/", func(c fiber.Ctx) error {
        var report models.Report
        if err := json.Unmarshal(c.Body(), &report); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to parse request body",
			})
		}
        report.Id = "12345" // Replace with actual logic to generate unique IDs
        return c.JSON(report)
    })

	// Retrieve report route (GET /api/v1/reports/:id)
	reports.Get("/:id", func(c fiber.Ctx) error {
        id := c.Params("id")
        // Replace this logic with actual database query or logic to retrieve the report
        report := models.Report{Id:id, Title: "Daily Report", Content: "This is a daily report", Type: "daily", SubmittedBy: "John Doe"}
        return c.JSON(report)
    })

    // Update report route (PUT /api/v1/reports/:id)
	reports.Put("/:id", func(c fiber.Ctx) error {
        id := c.Params("id")
        var report models.Report
        if err := json.Unmarshal(c.Body(), &report); err!= nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Failed to parse request body",
            })
        }
        report.Id = id
        // Replace this logic with actual database update or logic to update the report
        return c.JSON(report)
    })

	// List route (GET /api/v1/reports)
    // This route returns a list of all available reports in JSON format.
	reports.Get("/", func(c fiber.Ctx) error {
        reports := []models.Report{
            {Title: "Daily Report", Content: "This is a daily report", Type: "daily", SubmittedBy: "John Doe"},
            {Title: "Weekly Report", Content: "This is a weekly report", Type: "weekly", SubmittedBy: "Jane Doe"},
        }
        fmt.Println(reports);
        return c.JSON(reports) // Return a list of reports in JSON format     }
    })

    // Delete report route (DELETE /api/v1/reports/:id)
	reports.Delete("/:id", func(c fiber.Ctx) error {
        id := c.Params("id")
        // Replace this logic with actual database deletion or logic to delete the report
        return c.JSON(c.Status(200), "Report with ID "+id+" has been deleted")
    })

    // Search route (GET /api/v1/reports/search)
    // This route returns a list of reports based on the provided search query.
	reports.Get("/search", func(c fiber.Ctx) error {
        // query := c.Query("query")
        // Replace this logic with actual database query or logic to search for reports
        results := []models.Report{
            {Title: "Daily Report", Content: "This is a daily report", Type: "daily", SubmittedBy: "John Doe"},
            {Title: "Weekly Report", Content: "This is a weekly report", Type: "weekly", SubmittedBy: "Jane Doe"},
        }
        return c.JSON(results) // Return a list of matching reports in JSON format
    })

    // Download route (GET /api/v1/reports/:id/download)
    // This route returns a PDF file containing the report details.
	reports.Get("/:id/download", func(c fiber.Ctx) error {
        id := c.Params("id")
        // Replace this logic with actual database query or logic to retrieve the report details
        report := models.Report{Id: id, Title: "Daily Report", Content: "This is a daily report", Type: "daily", SubmittedBy: "John Doe"}
		fmt.Println(report);
        // Generate PDF file from report details and return it as a downloadable file
        return c.SendFile("./report.pdf") // Replace "./report.pdf" with actual file path
    })

    // Export route (GET /api/v1/reports/export)
    // This route exports all available reports in a specified format (e.g., CSV, JSON) and returns the exported file as a downloadable file.
	reports.Get("/export", func(c fiber.Ctx) error {
        // format := c.Query("format")
        // Replace this logic with actual database query or logic to retrieve all reports
        reports := []models.Report{
            {Title: "Daily Report", Content: "This is a daily report", Type: "daily", SubmittedBy: "John Doe"},
            {Title: "Weekly Report", Content: "This is a weekly report", Type: "weekly", SubmittedBy: "Jane Doe"},
        }
		fmt.Println(reports);
        // Export reports to specified format (e.g., CSV, JSON) and return it as a downloadable file
        return c.SendFile("./reports.csv") // Replace "./reports.csv" with actual file path
    })




	// Listening to port 3000
	log.Fatal(app.Listen(":3000"))
}
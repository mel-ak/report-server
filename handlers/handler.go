package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mel-ak/report-server/models"
)

func CreateReport(c *fiber.Ctx) error {
	var report models.Report
	if err := json.Unmarshal(c.Body(), &report); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	report.Id = "12345" // Replace with actual logic to generate unique IDs
	return c.JSON(report)
}

func GetReports(c *fiber.Ctx) error {
	reports := []models.Report{
		{Title: "Daily Report", Content: "This is a daily report", Type: "daily", SubmittedBy: "John Doe"},
		{Title: "Weekly Report", Content: "This is a weekly report", Type: "weekly", SubmittedBy: "Jane Doe"},
	}
	fmt.Println(reports)
	return c.JSON(reports) // Return a list of reports in JSON format     }
}

func GetReport(c *fiber.Ctx) error {
	id := c.Params("id")
	// Replace this logic with actual database query or logic to retrieve the report
	report := models.Report{Id: id, Title: "Daily Report", Content: "This is a daily report", Type: "daily", SubmittedBy: "John Doe"}
	return c.JSON(report)
}

func UpdateReport(c *fiber.Ctx) error {
	id := c.Params("id")
	var report models.Report
	if err := json.Unmarshal(c.Body(), &report); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	report.Id = id
	// Replace this logic with actual database update or logic to update the report
	return c.JSON(report)
}

func DeleteReport(c *fiber.Ctx) error {
	id := c.Params("id")
	// Replace this logic with actual database deletion or logic to delete the report
	return c.JSON(c.Status(200), "Report with ID "+id+" has been deleted")
}

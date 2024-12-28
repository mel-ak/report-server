package test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/mel-ak/report-server/database"
	"github.com/mel-ak/report-server/handlers"
	"github.com/mel-ak/report-server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setup() *fiber.App {
	app := fiber.New()
	database.DB, _ = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	database.DB.AutoMigrate(&models.Report{}) // Ensure the DB schema is set up
	return app
}

func TestCreateReport(t *testing.T) {
	app := setup()
	app.Post("/reports", handlers.CreateReport)

	report := models.Report{Title: "Test Report", Content: "This is a test report."}
	body, _ := json.Marshal(report)

	req := httptest.NewRequest("POST", "/reports", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json") // Set the Content-Type header

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		// Read the response body to get more information about the error
		var errorResponse fiber.Map
		json.NewDecoder(resp.Body).Decode(&errorResponse)
		t.Fatalf("Expected status %d, but got %d. Error: %v", fiber.StatusOK, resp.StatusCode, errorResponse)
	}
}
func TestGetReports(t *testing.T) {
	app := setup()
	app.Get("/reports", handlers.GetReports)

	// Add a report to test retrieval
	database.DB.Create(&models.Report{Title: "Test Report", Content: "This is a test report."})

	req := httptest.NewRequest("GET", "/reports", nil)
	resp, err := app.Test(req)

	if err != nil || resp.StatusCode != fiber.StatusOK {
		t.Fatalf("Expected status %d, but got %d", fiber.StatusOK, resp.StatusCode)
	}
}

func TestGetReport(t *testing.T) {
	app := setup()
	app.Get("/reports/:id", handlers.GetReport)

	report := models.Report{Title: "Test Report", Content: "This is a test report."}
	database.DB.Create(&report)

	req := httptest.NewRequest("GET", "/reports/"+strconv.FormatUint(uint64(report.ID), 10), nil)
	resp, err := app.Test(req)

	if err != nil || resp.StatusCode != fiber.StatusOK {
		t.Fatalf("Expected status %d, but got %d", fiber.StatusOK, resp.StatusCode)
	}
}

func TestUpdateReport(t *testing.T) {
	app := setup()
	app.Put("/reports/:id", handlers.UpdateReport)

	// Create a report to update
	report := models.Report{Title: "Test Report", Content: "This is a test report."}
	database.DB.Create(&report)

	// Prepare the updated report data
	updatedReport := models.Report{Title: "Updated Report", Content: "This is an updated report."}
	body, _ := json.Marshal(updatedReport)

	// Create the request for updating the report
	req := httptest.NewRequest("PUT", "/reports/"+strconv.FormatUint(uint64(report.ID), 10), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json") // Set Content-Type header

	// Send the request
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != fiber.StatusOK {
		// Read the response body to get more information about the error
		var errorResponse fiber.Map
		json.NewDecoder(resp.Body).Decode(&errorResponse)
		t.Fatalf("Expected status %d, but got %d. Error: %v", fiber.StatusOK, resp.StatusCode, errorResponse)
	}
}

func TestDeleteReport(t *testing.T) {
	app := setup()
	app.Delete("/reports/:id", handlers.DeleteReport)

	report := models.Report{Title: "Test Report", Content: "This is a test report."}
	database.DB.Create(&report)

	req := httptest.NewRequest("DELETE", "/reports/"+strconv.FormatUint(uint64(report.ID), 10), nil)
	resp, err := app.Test(req)

	if err != nil || resp.StatusCode != fiber.StatusNoContent {
		t.Fatalf("Expected status %d, but got %d", fiber.StatusNoContent, resp.StatusCode)
	}
}

// Benchmark test for GetReports
func BenchmarkGetReports(b *testing.B) {
	app := setup()
	app.Get("/reports", handlers.GetReports)

	// Populate the database with some test data
	for i := 0; i < 100; i++ {
		database.DB.Create(&models.Report{Title: "Test Report", Content: "This is a test report."})
	}

	b.ResetTimer() // Reset the timer to only measure the actual test

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/reports", nil)
		resp, err := app.Test(req)

		if err != nil {
			b.Fatalf("Failed to send request: %v", err)
		}

		if resp.StatusCode != fiber.StatusOK {
			b.Fatalf("Expected status %d, but got %d", fiber.StatusOK, resp.StatusCode)
		}
	}
}

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mel-ak/report-server/database"
	"github.com/mel-ak/report-server/models"
)

func CreateReport(c *fiber.Ctx) error {
	report := new(models.Report)
	if err := c.BodyParser(report); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Create(&report)
	return c.JSON(report)
}

func GetReports(c *fiber.Ctx) error {
	var reports []models.Report
	database.DB.Find(&reports)
	return c.JSON(reports)
}

func GetReport(c *fiber.Ctx) error {
	id := c.Params("id")
	var report models.Report
	if result := database.DB.First(&report, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Report not found"})
	}
	return c.JSON(report)
}

func UpdateReport(c *fiber.Ctx) error {
	id := c.Params("id")
	var report models.Report
	if result := database.DB.First(&report, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Report not found"})
	}
	if err := c.BodyParser(&report); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	database.DB.Save(&report)
	return c.JSON(report)
}

func DeleteReport(c *fiber.Ctx) error {
	id := c.Params("id")
	var report models.Report
	if result := database.DB.First(&report, id); result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Report not found"})
	}
	database.DB.Delete(&report)
	return c.SendStatus(fiber.StatusNoContent)
}

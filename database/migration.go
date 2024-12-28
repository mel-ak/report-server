package database

import "github.com/mel-ak/report-server/models"

func Migrate() {
	DB.AutoMigrate(&models.Report{})
	DB.AutoMigrate(&models.User{})
}

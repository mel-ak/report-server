package models

import "gorm.io/gorm"

type Report struct {
	gorm.Model
	Title       string `json:"title"`
	Content     string `json:"content"`
	Type        string `json:"type"` // "daily" or "weekly"
	SubmittedBy string `json:"submitted_by"`
}

package models

type Report struct {
	Id          uint64 `gorm:"auto_increment;column:id;type:bigint;primary_key"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Type        string `json:"type"` // "daily" or "weekly"
	SubmittedBy string `json:"submitted_by"`
}

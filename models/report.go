package models


type Report struct {
	Id          string `json:"id"`
    Title       string `json:"title"`
    Content     string `json:"content"`
    Type        string `json:"type"` // "daily" or "weekly"
    SubmittedBy string `json:"submitted_by"`
}
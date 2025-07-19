package models

type TaskDescription struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	DueDate     string   `json:"dueDate"`
	Priority    string   `json:"priority"`
	Tags        []string `json:"tags"`
}

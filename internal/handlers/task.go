package handlers

import (
	"calandar-desktop-task/internal/models"
	"time"
)

func CreateNewTask(Description string) *models.TaskDescription {
	return &models.TaskDescription{
		Title:       Description,
		Description: Description,
		DueDate:     time.Now().AddDate(0, 0, 1).Format(time.RFC3339),
		Priority:    "", // low, medium, high
		Tags:        []string{},
	}
}

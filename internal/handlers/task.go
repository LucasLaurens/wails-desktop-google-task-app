package handlers

import "calandar-desktop-task/internal/models"

func CreateNewTask() *models.Task {
	Description := "My first little task"

	return &models.Task{
		Description: Description,
	}
}

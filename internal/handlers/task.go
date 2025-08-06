package handlers

import (
	"log"
	"time"

	"google.golang.org/api/tasks/v1"
)

func CreateNewTask(Title string) *tasks.Task {
	// todo: replace by env var
	loc, err := time.LoadLocation("Europe/Paris")
	if err != nil {
		log.Fatalf("failed to load location: %v", err)
	}

	nowParis := time.Now().AddDate(0, 0, 7).In(loc)

	return &tasks.Task{
		Title: Title,
		Due:   nowParis.Format(time.RFC3339),
	}
}

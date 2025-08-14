package handlers

import (
	"calandar-desktop-task/internal/errors"
	"os"
	"time"

	"google.golang.org/api/tasks/v1"
)

func CreateNewTask(Title string) *tasks.Task {
	location := getLocation()
	now := time.Now().AddDate(0, 0, 7).In(location)

	return &tasks.Task{
		Title: Title,
		Due:   now.Format(time.RFC3339),
	}
}

func getLocation() *time.Location {
	defaultTimezone := os.Getenv("DEFAULT_TIMEZONE")
	location, err := time.LoadLocation(defaultTimezone)
	errors.Fatal(
		"failed to load location by '%v' default timezone with the '%v' error",
		errors.FatalError{
			Err:  err,
			Args: []interface{}{defaultTimezone},
		},
	)

	return location
}

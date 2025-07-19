package main

import (
	"calandar-desktop-task/internal/handlers"
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.RegisterTaskCreationListener()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) RegisterTaskCreationListener() {
	runtime.EventsOn(
		a.ctx,
		"createTask",
		func(args ...interface{}) {
			if len(args) == 0 {
				fmt.Println("No arguments received...")
				return
			}

			if _, ok := args[0].(string); !ok {
				fmt.Println("Received task is not a string")
				return
			}

			fmt.Printf(
				"The new task: %v \n",
				args[0],
			)
		},
	)
}

// todo: replace the sync call by event emit/on
func (a *App) WailsEmitTaskWithDescription() string {
	task := handlers.CreateNewTask()

	return task.Description
}

// func (a *App) WailsAsyncTask() error {
// newChan := make(chan struct{})

// go func() {
// 	runtime.EventsEmit(
// 		a.ctx,
// 		"createNewTask",
// 		"qsdsqdsd",
// 	)
// close(newChan)
// }()
// <-newChan
// }

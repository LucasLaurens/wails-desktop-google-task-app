package main

import (
	api "calandar-desktop-task/external/api/google"
	"calandar-desktop-task/internal/handlers"
	"context"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
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
	taskService, err := a.RegisterGoogleTaskServiceProvider()
	if err != nil {
		fmt.Println(err)
	}

	// test to get the first ten tasks
	taskService.GetTasksList(10)

	a.RegisterTaskCreationListener()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) RegisterGoogleTaskServiceProvider() (api.TaskServiceWrapper, error) {
	// todo: replace by env var
	readByte, err := os.ReadFile("external/api/google/credentials.json")
	if err != nil {
		return api.TaskServiceWrapper{}, fmt.Errorf(
			"unable to read client secret file: %v",
			err,
		)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(readByte, tasks.TasksReadonlyScope)
	if err != nil {
		return api.TaskServiceWrapper{}, fmt.Errorf(
			"unable to parse client secret file to config: %v",
			err,
		)
	}

	client := api.GetClient(config)
	taskService, err := tasks.NewService(
		a.ctx,
		option.WithHTTPClient(client),
	)
	if err != nil {
		return api.TaskServiceWrapper{}, fmt.Errorf(
			"unable to retrieve tasks Client: %v",
			err,
		)
	}

	return api.TaskServiceWrapper{
		Service: taskService,
	}, nil
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

			if task, ok := args[0].(string); ok {
				fmt.Println(task)
				newTask := handlers.CreateNewTask(task)
				fmt.Printf(
					"The new task: %v for %v \n",
					newTask.Title,
					newTask.DueDate,
				)

				return
			}

			fmt.Println("Received task is not a string")
		},
	)
}

package main

import (
	api "calandar-desktop-task/external/api/google"
	"calandar-desktop-task/internal/handlers"
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/tasks/v1"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	godotenv.Load()

	a.ctx = ctx
	taskService, err := a.RegisterGoogleTaskServiceProvider()
	if err != nil {
		fmt.Println(err)
	}

	// test to get the first ten tasks
	taskService.GetTasksList(10)

	a.RegisterTaskCreationListener(taskService)
}

func (a *App) RegisterGoogleTaskServiceProvider() (api.TaskServiceWrapper, error) {
	readByte, err := os.ReadFile(os.Getenv("CREDENTIALS_PATH"))
	if err != nil {
		return api.TaskServiceWrapper{}, fmt.Errorf(
			"unable to read client secret file: %v",
			err,
		)
	}

	config, err := google.ConfigFromJSON(readByte, tasks.TasksScope)
	if err != nil {
		return api.TaskServiceWrapper{}, fmt.Errorf(
			"unable to parse client secret file to config: %v",
			err,
		)
	}

	client := api.GetClient(a.ctx, config)
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

func (a *App) RegisterTaskCreationListener(taskService api.TaskServiceWrapper) {
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
					newTask.Due,
				)

				taskService.InsertNewTask(newTask)
				taskService.GetTasksList(10)

				return
			}

			fmt.Println("Received task is not a string")
		},
	)
}

package api

import (
	internalConfig "calandar-desktop-task/internal/config"
	"calandar-desktop-task/internal/errors"
	"calandar-desktop-task/internal/server"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"google.golang.org/api/tasks/v1"
)

type TaskServiceWrapper struct {
	Service *tasks.Service
}

type TaskList struct {
	*tasks.Tasks
}

func (taskService *TaskServiceWrapper) InsertNewTask(task *tasks.Task) {
	id := internalConfig.GetConfig("TASK_LIST_ID")
	_, err := taskService.Service.Tasks.Insert(id, task).Do()

	errors.Fatal(
		"unable to insert a new task: %v",
		errors.FatalError{
			Err:  err,
			Args: []interface{}{},
		},
	)
}

func (taskService *TaskServiceWrapper) GetTasksList(max int64) {
	id := internalConfig.GetConfig("TASK_LIST_ID")
	list, err := taskService.Service.Tasks.List(id).Do()

	errors.Fatal(
		"unable to retrieve task lists. %v",
		errors.FatalError{
			Err:  err,
			Args: []interface{}{},
		},
	)

	taskList := TaskList{list}
	taskList.displayTaskListItems()
}

func (taskList *TaskList) displayTaskListItems() {
	fmt.Println("Task Lists:")
	if len(taskList.Items) <= 0 {
		fmt.Println("no tasks ar available")
	}

	for _, item := range taskList.Items {
		fmt.Printf("%s (%s)\n", item.Title, item.Status)
	}
}

func GetClient(ctx context.Context, config *oauth2.Config) *http.Client {
	newTokenFileName := internalConfig.GetConfig("TOKEN_PATH")
	token, err := newTokenFromFile(newTokenFileName)

	if err != nil || !token.Valid() {
		token := getTokenFromWeb(ctx, config)
		saveToken(newTokenFileName, token)
	}

	return config.Client(context.Background(), token)
}

func newTokenFromFile(tokenFileName string) (*oauth2.Token, error) {
	file, err := os.Open(tokenFileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)
	return token, err
}

func getTokenFromWeb(ctx context.Context, config *oauth2.Config) *oauth2.Token {
	token := server.Handle(config, ctx)

	return token
}

func saveToken(tokenFileName string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", tokenFileName)

	file, err := os.OpenFile(
		tokenFileName,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0600,
	)

	errors.Fatal(
		"Unable to cache oauth token: %v",
		errors.FatalError{
			Err:  err,
			Args: []interface{}{},
		},
	)

	defer file.Close()

	if token != nil {
		json.NewEncoder(file).Encode(token)
	}
}

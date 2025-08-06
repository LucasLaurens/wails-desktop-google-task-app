package api

import (
	"calandar-desktop-task/internal/server"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"google.golang.org/api/tasks/v1"
)

// todo: move to specific struct file
type TaskServiceWrapper struct {
	Service *tasks.Service
}

func (taskService *TaskServiceWrapper) InsertNewTask(task *tasks.Task) {
	// todo: declare as global
	// todo: fix getting env
	id := os.Getenv("TASK_LIST_ID")
	if id == "" {
		id = "MTMxMzU1MTg2Nzk4NzI1MTc0MTg6MDow"
	}

	_, err := taskService.Service.Tasks.Insert(id, task).Do()

	if err != nil {
		log.Fatalf("unable to insert a new task: %v", err)
	}
}

func (taskService *TaskServiceWrapper) GetTasksList(max int64) {
	// todo: create an oauth provider
	// todo: only using env var
	id := os.Getenv("TASK_LIST_ID")
	if id == "" {
		id = "MTMxMzU1MTg2Nzk4NzI1MTc0MTg6MDow"
	}

	list, err := taskService.Service.Tasks.List(id).Do()
	if err != nil {
		log.Fatalf("unable to retrieve task lists. %v", err)
	}

	fmt.Println("Task Lists:")
	if len(list.Items) <= 0 {
		fmt.Println("no tasks ar available")
	}

	for _, item := range list.Items {
		fmt.Printf("%s (%s)\n", item.Title, item.Status)
	}
}

func GetClient(ctx context.Context, config *oauth2.Config) *http.Client {
	// todo: only using env var
	newTokenFileName := "external/api/google/token.json"
	token, err := newTokenFromFile(newTokenFileName)
	if err != nil {
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
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}

	defer file.Close()

	if token != nil {
		json.NewEncoder(file).Encode(token)
	}
}

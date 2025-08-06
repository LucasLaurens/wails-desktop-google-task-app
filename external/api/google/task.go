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

func (taskService *TaskServiceWrapper) GetTasksList(max int64) error {
	id := "MTMxMzU1MTg2Nzk4NzI1MTc0MTg6MDow"
	list, err := taskService.Service.Tasks.List(id).Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve task lists. %v", err)
	}

	fmt.Println("Task Lists:")
	if len(list.Items) <= 0 {
		return fmt.Errorf("no tasks ar available")
	}

	for _, item := range list.Items {
		fmt.Printf("%s (%s)\n", item.Title, item.Status)
	}
	return nil
}

// Retrieve a token, saves the token, then returns the generated client.
func GetClient(ctx context.Context, config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	newTokenFileName := "external/api/google/token.json"
	token, err := newTokenFromFile(newTokenFileName)
	if err != nil {
		// todo: try context.Background()
		token := getTokenFromWeb(ctx, config)
		saveToken(newTokenFileName, token)
	}

	return config.Client(context.Background(), token)
}

// Retrieves a token from a local file.
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

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(ctx context.Context, config *oauth2.Config) *oauth2.Token {
	token := server.Handle(config, ctx)

	return token
}

// Saves a token to a file path.
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
		// store web token into an internal file
		json.NewEncoder(file).Encode(token)
	}
}

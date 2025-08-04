package api

import (
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
	list, err := taskService.Service.Tasklists.List().MaxResults(max).Do()
	if err != nil {
		return fmt.Errorf("unable to retrieve task lists. %v", err)
	}

	fmt.Println("Task Lists:")
	if len(list.Items) <= 0 {
		return fmt.Errorf("no tasks ar available")
	}

	for _, item := range list.Items {
		fmt.Printf("%s (%s)\n", item.Title, item.Id)
	}
	return nil
}

// Retrieve a token, saves the token, then returns the generated client.
func GetClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	newTokenFileName := "external/api/google/token.json"
	token, err := newTokenFromFile(newTokenFileName)
	if err != nil {
		token = getTokenFromWeb(config)
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
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf(
		"Go to the following link in your browser then type the authorization code: %v \n",
		authURL,
	)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		fmt.Printf(
			"Unable to read authorization code: %v \n",
			err,
		)
	}

	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		fmt.Printf(
			"Unable to retrieve token from web: %v \n",
			err,
		)

	}

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

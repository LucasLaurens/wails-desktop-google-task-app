package config

import (
	"calandar-desktop-task/internal/errors"
	"os"
	"reflect"
)

type GoogleConfig struct {
	TASK_LIST_ID string
}

type Config struct {
	Domain          string
	Port            string
	BaseUrl         string
	DefaultTimezone string
	CredentialsPath string
	TokenPath       string
	Service         GoogleConfig
}

func GetConfig(key string) string {
	value := reflect.ValueOf(Config{})
	field := value.FieldByName(key)
	if !field.IsValid() {
		errors.Fatal(
			"The current variable does not exist from the configuraiton",
			errors.FatalError{},
		)
	}

	return os.Getenv(key)
}

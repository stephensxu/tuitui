package tuitui

import (
	"fmt"
	"os/user"
	"path/filepath"
	"io/ioutil"
	"github.com/segmentio/go-prompt"
	"encoding/json"
)


func NewAuthenticatedClient() (*TuituiClient, error) {
	client := NewTuituiClient()
	authenticated := client.Authenticate()

	if authenticated {
		return client, nil
	}

	return nil, fmt.Errorf("Could not load authenticated client")
}

func Login(client *TuituiClient) (bool, error) {
	fields := map[string]string{}

	for k, v := range GetAuthFields() {
		message := fmt.Sprintf("Please enter your twitter %v", k)
		promptValue := ""
		if v {
			promptValue = prompt.PasswordMasked(message)
		} else {
			promptValue = prompt.String(message)
		}

		fields[k] = promptValue
	}

	err := Save(fields)

	if err != nil {
		return false, err
	}

	authenticated := client.Authenticate()

	if authenticated {
		return true, nil
	}

	return false, fmt.Errorf("Could not login")
}

func Load() (map[string]string, error) {
	fileLocation := getFileLocation()
	var c = map[string]string{}

	file, err := ioutil.ReadFile(fileLocation)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &c)

	if err != nil {
		return nil, err
	}

	return c, nil
}

// Save : saves the configuration to a file
func Save(values map[string]string) error {
	fileLocation := getFileLocation()
	loginDetails, err := json.Marshal(values)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileLocation, loginDetails, 0644)

	if err != nil {
		return err
	}

	return nil
}

func getFileLocation() string {
	dir := getUserHomeOrDefault()
	return filepath.Join(dir, ".tuitui.json")
}

func getUserHomeOrDefault() string {
	usr, err := user.Current()

	if err != nil {
		return "./"
	}

	return usr.HomeDir
}

func GetAuthFields() map[string]bool {
	return map[string]bool{
		"consumerKey": true,
		"consumerSecret": true,
		"accessToken": true,
		"accessSecret": true,
	}
}

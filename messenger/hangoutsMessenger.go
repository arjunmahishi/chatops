package messenger

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/arjunmahishi/Chatops/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// HangoutsMessenger implements Messenger for Hangouts chat
type HangoutsMessenger struct {
	Client *http.Client
}

func (hm *HangoutsMessenger) setup() error {
	credsFile, err := os.Open(config.Config.ServiceAccountCredsPath)
	if err != nil {
		return err
	}
	defer credsFile.Close()

	credsJSON, err := ioutil.ReadAll(credsFile)
	if err != nil {
		return err
	}

	config, err := google.JWTConfigFromJSON(credsJSON, "https://www.googleapis.com/auth/chat.bot")
	if err != nil {
		return err
	}
	hm.Client = config.Client(oauth2.NoContext)
	return nil
}

// Send a message to the given space
func (hm *HangoutsMessenger) Send(space string, message string) error {
	messageReader := strings.NewReader(fmt.Sprintf(`{"text": "%s"}`, strings.Replace(message, "\"", "'", -1)))
	req, err := http.NewRequest("POST", "https://chat.googleapis.com/v1/"+space+"/messages", messageReader)
	if err != nil {
		return err
	}
	_, err = hm.Client.Do(req) // Handle 200 errors
	if err != nil {
		return err
	}
	return nil
}

// Update a message with a new message
func (hm *HangoutsMessenger) Update(space, messageName, newMessage string) error {
	messageReader := strings.NewReader("{text: \"" + newMessage + "\"}")
	req, err := http.NewRequest("PUT", "https://chat.googleapis.com/v1/"+space+"/messages/"+messageName+"?updateMask=text", messageReader)
	if err != nil {
		return err
	}

	res, err := hm.Client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return nil
}

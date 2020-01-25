package payload

import (
	"regexp"
	"strings"
)

// HangoutsPayload structure
type HangoutsPayload struct {
	Type      string
	EventTime string
	Token     string
	Message   message
	User      user
	Space     space
}

type message struct {
	Name         string
	ArgumentText string
}

type user struct {
	Name        string
	DisplayName string
	Email       string
	Type        string
}

type space struct {
	Name        string
	Type        string
	DisplayName string
}

// Validate the recived payload
func (hp HangoutsPayload) Validate() bool {
	if !validateTime(hp.EventTime) {
		return false
	}
	return true
}

// GetMessage from payload
func (hp HangoutsPayload) GetMessage() string {
	message := strings.Trim(hp.Message.ArgumentText, " ")
	return removeMentions(message)
}

// GetSenderName from payload
func (hp HangoutsPayload) GetSenderName() string {
	return hp.User.DisplayName
}

// GetSenderEmail from payload
func (hp HangoutsPayload) GetSenderEmail() string {
	return hp.User.Email
}

// GetSpace name of the current conversation
func (hp HangoutsPayload) GetSpace() string {
	return hp.Space.Name
}

func removeMentions(text string) string {
	re := regexp.MustCompile(`@\S+`)
	text = re.ReplaceAllString(text, "")
	return strings.TrimSpace(text)
}

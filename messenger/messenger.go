package messenger

import "log"

// Messenger interface to send or update messages to a user/room
type Messenger interface {
	setup() error
	Send(space string, message string) error
	Update(space, messageName, newMessage string) error
}

// NewMessenger creates a new instance of Messenger
func NewMessenger() Messenger {
	messenger := &HangoutsMessenger{}
	err := messenger.setup()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return messenger
}

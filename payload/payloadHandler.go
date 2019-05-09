package payload

import (
	"fmt"
	"time"
)

//Handler interface
type Handler interface {
	GetMessage() string
	Validate() bool
	GetSenderName() string
	GetSenderEmail() string
	GetSpace() string
}

func validateTime(eventTime string) bool {
	currentTime := time.Now().Add(-1 * time.Minute)

	requestTime, err := time.Parse(time.RFC3339, eventTime)
	if err != nil {
		return false
	}
	return requestTime.After(currentTime)
}

// CreatePayloadHandler initializes a new PayloadHandler for
// a given backend choice
func CreatePayloadHandler(backend string) (Handler, error) {
	switch backend {
	case "hangouts":
		var payload HangoutsPayload
		return &payload, nil
	default:
		return nil, fmt.Errorf("couldn't find a backend called %s", backend)
	}
}

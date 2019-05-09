package commanders

import (
	"regexp"

	"github.com/arjunmahishi/Chatops/payload"
)

// Commander interface
type Commander interface {
	GetName() string
	GetCatagory() string
	MatchCommand(string) bool
	Execute(payload.Handler) (map[string]interface{}, error)
}

// GetCommander returns an appropriate Commander by matching the text with existing templates
func GetCommander(text string) (Commander, error) {
	for _, commander := range CommanderList {
		if commander.MatchCommand(text) {
			return commander, nil
		}
	}

	return DialogFlowCommand{
		Name:         "small-talk",
		Catagory:     "dialogflow",
		OutputFormat: "",
		RegexPattern: regexp.MustCompile(`(?mi)(wfh|ooo)\s+(\d{4}-\d{1,2}-\d{1,2})`),
	}, nil
}

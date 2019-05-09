package commanders

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/arjunmahishi/Chatops/aggregators"
	"github.com/arjunmahishi/Chatops/messenger"
	"github.com/arjunmahishi/Chatops/payload"
)

// LocalBashCommand structure
type LocalBashCommand struct {
	Command      string
	Name         string
	Catagory     string
	OutputFormat string
	Scope        int64
	RegexPattern *regexp.Regexp
}

// Execute given command
func (bc LocalBashCommand) Execute(payloadHandler payload.Handler) (map[string]interface{}, error) {
	bc.reconstructCommand(payloadHandler.GetMessage())

	go func() {
		sender := messenger.NewMessenger()
		output, err := runCommandLocally(bc.Command)
		if err != nil {
			log.Printf(err.Error())
			sender.Send(payloadHandler.GetSpace(), fmt.Sprintf("Sorry your request returned an error:\n```%s```", err.Error()))
			return
		}

		aggregator := aggregators.GetAggregator(string(output), bc.OutputFormat)
		err = aggregator.Parse()
		if err != nil {
			log.Printf(err.Error())
			sender.Send(payloadHandler.GetSpace(), fmt.Sprintf("Sorry your request returned an error:\n```%s```", err.Error()))
			return
		}

		log.Printf("Sent the output of '%s' to %s(%s)", bc.Name, payloadHandler.GetSenderName(), payloadHandler.GetSenderEmail())
		sender.Send(payloadHandler.GetSpace(), fmt.Sprintf("Here is your result for the request `%s`:\n%s", bc.Name, aggregator.String()))
	}()
	return map[string]interface{}{"text": fmt.Sprintf("Your request for `%s` is being processed", bc.Name)}, nil
}

// MatchCommand from the list of commands
func (bc LocalBashCommand) MatchCommand(text string) bool {
	return bc.RegexPattern.MatchString(text)
}

// GetName the name of the command that was execute
func (bc LocalBashCommand) GetName() string {
	return bc.Name
}

// GetCatagory that was executed
func (bc LocalBashCommand) GetCatagory() string {
	return bc.Catagory
}

// GetScope that is required to run this command
func (bc LocalBashCommand) GetScope() int64 {
	return bc.Scope
}

func (bc *LocalBashCommand) reconstructCommand(chatText string) {
	match := bc.RegexPattern.FindStringSubmatch(chatText)
	result := make(map[string]string)

	for i, name := range bc.RegexPattern.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}

	for mark, value := range result {
		bc.Command = strings.Replace(bc.Command, fmt.Sprintf("{{%s}}", mark), value, -1)
	}
}

func runCommandLocally(command string) ([]byte, error) {
	log.Printf("Executing command [ %s ]", command)
	out, err := exec.Command("/bin/sh", "-c", command).Output()
	if err != nil {
		return nil, err
	}

	out = bytes.TrimSpace(out)
	out = bytes.Trim(out, "\n")
	return out, nil
}

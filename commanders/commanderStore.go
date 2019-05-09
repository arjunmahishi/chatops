package commanders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
)

// CommanderList contains a list of available commands
var CommanderList []Commander

// CommandsData holds the commands in the json format read from commands.json
var CommandsData struct {
	Commands []struct {
		Name     string
		Hostname string
		Command  string
		Regex    string
		Example  string
	}
}

var staticCommanderList = []Commander{}

// SyncCommands with the database
func SyncCommands(commandsPath string) error {
	CommanderList = staticCommanderList
	conts, err := ioutil.ReadFile(commandsPath)
	if err != nil {
		return err
	}
	json.Unmarshal(conts, &CommandsData)

	log.Printf("Syncing commands list")
	for _, command := range CommandsData.Commands {
		if command.Hostname != "" {
			command.Command = fmt.Sprintf("ssh %s %s", command.Hostname, command.Command)
		}

		var lCommand LocalBashCommand
		lCommand.Name = command.Name
		lCommand.Command = command.Command
		lCommand.Catagory = "bash"
		lCommand.OutputFormat = ""
		lCommand.RegexPattern = regexp.MustCompile(fmt.Sprintf(`(?mi)%s`, command.Regex))
		CommanderList = append(CommanderList, lCommand)
	}

	log.Printf("Total commands %d", len(CommanderList))
	return nil
}

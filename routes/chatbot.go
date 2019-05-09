package routes

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/arjunmahishi/Chatops/commanders"
	"github.com/arjunmahishi/Chatops/payload"
	echo "github.com/labstack/echo"
)

var getCommander = commanders.GetCommander

// Chatbot endpoint
func Chatbot(c echo.Context) error {
	payload, err := payload.CreatePayloadHandler("hangouts")
	if err != nil {
		panic(err)
	}
	c.Bind(&payload)

	if !payload.Validate() {
		return c.String(http.StatusForbidden, "Not allowed")
	}
	return c.JSON(http.StatusOK, getResponse(payload))
}

func getResponse(payloadHandler payload.Handler) map[string]interface{} {
	var response = map[string]interface{}{"text": ""}

	if strings.ToLower(payloadHandler.GetMessage()) == "help" {
		response["text"] = getHelpText()
		return response
	}

	commander, err := getCommander(payloadHandler.GetMessage())
	if err != nil {
		response["text"] = err.Error()
		return response
	}

	// Audit-log
	logData(payloadHandler, commander)
	return execute(commander, payloadHandler)
}

func execute(commander commanders.Commander, payload payload.Handler) map[string]interface{} {
	output, err := commander.Execute(payload)
	if err != nil {
		return map[string]interface{}{"text": err.Error()}
	}

	return output
}

func logData(payloadHandler payload.Handler, commander commanders.Commander) {
	log.Printf("Running command %s for %s(%s)", commander.GetName(), payloadHandler.GetSenderName(), payloadHandler.GetSenderEmail())
}

func getHelpText() string {
	helpText := "*Here are a list of commands that chatops can run:*\n"
	for i, command := range commanders.CommandsData.Commands {
		helpText += fmt.Sprintf("%d) %s \n\t\t*Ex:* `%s`\n", i+1, command.Name, command.Example)
	}
	return helpText
}

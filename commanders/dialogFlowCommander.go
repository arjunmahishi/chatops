package commanders

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/arjunmahishi/Chatops/config"
	"github.com/arjunmahishi/Chatops/payload"
)

var accessToken = config.Config.DialogFlowAccessToken

// DialogFlowCommand structure
type DialogFlowCommand struct {
	Name         string
	Catagory     string
	OutputFormat string
	Scope        int64
	RegexPattern *regexp.Regexp
}

type dialogFlowResponse struct {
	Result struct {
		Action      string `json:"action"`
		Fulfillment struct {
			Speech string `json:"speech"`
		}
	} `json:"result"`
}

// GetName of the given command
func (uc DialogFlowCommand) GetName() string {
	return uc.Name
}

// GetCatagory of the given command
func (uc DialogFlowCommand) GetCatagory() string {
	return uc.Catagory
}

// GetScope required to execute this command
func (uc DialogFlowCommand) GetScope() int64 {
	return uc.Scope
}

// MatchCommand based of regex
func (uc DialogFlowCommand) MatchCommand(text string) bool {
	return uc.RegexPattern.MatchString(text)
}

// Execute the given command
func (uc DialogFlowCommand) Execute(payload payload.Handler) (map[string]interface{}, error) {
	dialogResponse, err := sendDialog(payload.GetMessage())
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{"text": dialogResponse.Result.Fulfillment.Speech}, nil
}

func sendDialog(query string) (*dialogFlowResponse, error) {
	data := map[string]string{
		"query":     query,
		"lang":      "en",
		"sessionId": "sessionId",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.api.ai/v1/query?v=20150910", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var dRes dialogFlowResponse
	err = json.Unmarshal(resBody, &dRes)

	return &dRes, nil
}

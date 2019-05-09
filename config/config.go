package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

// Config values that are read from the config file
var Config struct {
	BotName                 string
	HangoutsToken           string
	DialogFlowAccessToken   string
	CommandsPath            string
	ServiceAccountCredsPath string
}

var configPath = flag.String("config", "config.json", "path to config.json file")

func init() {
	flag.Parse()
	if err := populateConfig(*configPath); err != nil {
		log.Fatalf("Couldn't read the config file: %s", err.Error())
	}
}

func populateConfig(configPath string) error {
	conts, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(conts, &Config)
}

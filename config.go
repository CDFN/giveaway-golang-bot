package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func loadConfig(){
	var config Config
	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal([]byte(configFile), &config)
	if err != nil {
		panic(err)
	}

	Token = config.Token
	Prefix = config.Prefix
	GiveawayHours = config.Hours
	GiveawayMinutes = config.Minutes
	GiveawaySeconds = config.Seconds
	AcceptingEnabled = config.AcceptingEnabled
	RoleNeededToAccept = config.RoleNeededToAccept

}
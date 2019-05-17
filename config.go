package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct{
	Prefix  string `yaml:"prefix"`
	Hours   int    `yaml:"giveawayHours"`
	Minutes int    `yaml:"giveawayMinutes"`
	Seconds int    `yaml:"giveawaySeconds"`
	Token   string `yaml:"botToken"`
}

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
}
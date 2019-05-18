package main

type Config struct{
	Prefix  string `yaml:"prefix"`
	Hours   int    `yaml:"giveawayHours"`
	Minutes int    `yaml:"giveawayMinutes"`
	Seconds int    `yaml:"giveawaySeconds"`
	Token            string `yaml:"botToken"`
	AcceptingEnabled bool `yaml:"acceptingEnabled"`
	RoleNeededToAccept string `yaml:"roleNeededToAccept"`
}


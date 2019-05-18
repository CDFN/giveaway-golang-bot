package main

import "github.com/bwmarrin/discordgo"

type Config struct{
	Prefix             string `yaml:"prefix"`
	Hours              int    `yaml:"giveawayHours"`
	Minutes            int    `yaml:"giveawayMinutes"`
	Seconds            int    `yaml:"giveawaySeconds"`
	Token              string `yaml:"botToken"`
	AcceptingEnabled   bool   `yaml:"acceptingEnabled"`
	RoleNeededToAccept string `yaml:"roleNeededToAccept"`
}

type Thx struct{
	messageID string
	status    Status
	mentioner *discordgo.User
	mentioned *discordgo.User
}
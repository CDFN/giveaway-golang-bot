package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == "" {
		return
	}
	if message.Author.ID == session.State.User.ID {
		return
	}

	args := strings.Fields(message.Content)

	if strings.HasPrefix(message.Content, Prefix) {

		switch strings.Trim(strings.ToLower(args[0]), "!") {
		case "help", "pomoc":
			sendMessage(session, message.ChannelID, "!thx @Mention - Say `thank you` and add mentioned to giveaway. You cant add yourself")
		case "thx", "thanks", "podziekuj":
			var mention *discordgo.User

			if len(message.Mentions) == 0 || len(args) < 2 {
				sendMessage(session, message.ChannelID, "Mention who you want thank to!")
				return
			} else {
				mention = message.Mentions[0]
			}
			if mention.ID == message.Author.ID {
				sendMessage(session, message.ChannelID, "Dont be like that! You cant thank yourself \\:)")
				return
			}
			Participants = append(Participants, mention.Username)
			sendMessage(session, message.ChannelID, "Participants: "+strings.Join(Participants, ", "))

		}

	}
}

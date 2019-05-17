package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == ""{
		return
	}
	if message.Author.ID == session.State.User.ID{
		return
	}

	args := strings.Fields(message.Content)

	if strings.HasPrefix(message.Content, Prefix){

		switch strings.Trim(strings.ToLower(args[0]), "!") {
		case "ping":
			sendMessage(session, message.ChannelID, ":ping_pong: Pong! ")
		case "reverse", "odwroc":
			if len(args)<2{
				return
			}
			s := reverse(strings.Join(args[1:], " "))
			sendMessage(session, message.ChannelID, s)
		case "thx", "thanks", "podziekuj":

			var mention *discordgo.User

			if len(message.Mentions) == 0 || len(args) < 2{
				sendMessage(session, message.ChannelID, "Musisz oznaczyć osobę, której chcesz podziękować!")
				return
			}else{
				mention = message.Mentions[0]
			}
			if mention.ID == message.Author.ID{
				sendMessage(session, message.ChannelID, "Nie bądź samolubny!  Nie możesz podziękowac samemu sobie \\:)")
				return
			}
			Participants = append(Participants, mention.Username)
			sendMessage(session, message.ChannelID, "Uczestnicy: " + strings.Join(Participants, ", "))




		default:
			sendMessage(session, message.ChannelID, "Nie znam komendy `tu bedzie komenda`!")

		}

	}
}

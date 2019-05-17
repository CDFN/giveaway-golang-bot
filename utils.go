package main

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func sendMessage(session *discordgo.Session, channelID string, message string) {
	_, err := session.ChannelMessageSend(channelID, message)
	checkErr("Błąd podczas wysyłania wiadomości", err)
}

func resolveGiveaway() {
	if !(len(Participants) > 0) {
		sendMessage(Session, "578612032052396052", "Brak uczestników!")
		return
	}
	randomNumber := rand.Intn(len(Participants))
	sendMessage(Session, "578612032052396052", "Wygrany: "+Participants[randomNumber])
	Participants = make([]string, 0)
}

func checkErr(msg string, err error) {
	if err != nil {
		fmt.Println(msg+":\n", err)
	}
}

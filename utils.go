package main

import (
	"fmt"
	"math/rand"

	"github.com/bwmarrin/discordgo"
)

func sendMessage(session *discordgo.Session, channelID string, message string) string {
	messageCreate, err := session.ChannelMessageSend(channelID, message)
	checkErr("Błąd podczas wysyłania wiadomości", err)
	return messageCreate.ID
	}

func resolveGiveaway() {
	Participants := getAllConfirmedUsers(Thxs)

	if !(len(Participants) > 0) {
		sendMessage(Session, "578612032052396052", "Brak uczestników!")
		return
	}

	randomNumber := rand.Intn(len(Participants))
	sendMessage(Session, "578612032052396052", "Wygrany: "+Participants[randomNumber].Username)
	Thxs = make([]Thx, 0)
}

func getAllUsernames(usersArray []*discordgo.User) []string{
	usernamesArray := make([]string, 0)
	for i := range usersArray{
		usernamesArray = append(usernamesArray, usersArray[i].Username)
	}
	return usernamesArray
}
func getAllConfirmedUsers(thxsArray []Thx) []*discordgo.User{
	usersArray := make([]*discordgo.User, 0)
	for i, participant := range Thxs{
		if participant.status == Accepted{
			usersArray = append(usersArray, thxsArray[i].mentioned)
		}
	}
	return usersArray
}

func checkErr(msg string, err error) {
	if err != nil {
		fmt.Println(msg+":\n", err)
	}
}

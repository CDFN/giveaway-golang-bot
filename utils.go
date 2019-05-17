package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
)


func sendMessage(session *discordgo.Session, channelID string, message string){
	_, err := session.ChannelMessageSend(channelID, message)
	checkErr("Błąd podczas wysyłania wiadomości", err)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func resolveGiveaway(){
	if !(len(Participants) > 0){
		sendMessage(Session, "578612032052396052", "Brak uczestników!")
		return
	}
	randomNumber := rand.Intn(len(Participants))
	sendMessage(Session, "578612032052396052", "Wygrany: " + Participants[randomNumber])
	Participants = make([]string, 0)
}
func checkErr(msg string, err error){
	if err != nil{
		fmt.Println(msg+":\n", err)
	}
}
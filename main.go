package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

// Variables used for command line parameters
var (
	Session            *discordgo.Session
	Token              string
	Prefix             string
	GiveawayHours      int
	GiveawayMinutes    int
	GiveawaySeconds    int
	AcceptingEnabled   bool
	RoleNeededToAccept string
	Thxs               []Thx
	)


func main() {
	loadConfig()
	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	checkErr("Error while creating session", err)
	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)
	dg.AddHandler(botEnable)
	dg.AddHandler(reactionAdd)
	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	checkErr("Error while opening connection", err)
	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.\nPress CTRL-C to exit.")


	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	err = dg.Close()
	checkErr("Error while closing session", err)
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.







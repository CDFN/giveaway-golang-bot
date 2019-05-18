package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/robfig/cron"
)

func botEnable(session *discordgo.Session, connect *discordgo.Ready){
	Session = session
	c := cron.New()
	err := c.AddFunc(fmt.Sprintf("%d %d %d * * *", GiveawaySeconds, GiveawayMinutes, GiveawayHours), resolveGiveaway)
	checkErr("Error while adding function to CRON!", err)
	c.Start()
}

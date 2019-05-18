package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)
//❌ ✅
func reactionAdd(session *discordgo.Session, reaction *discordgo.MessageReactionAdd){
	for i := 0; i < len(Thxs); i++{
		if reaction.MessageID == Thxs[i].messageID{
			member, err := session.State.Member(reaction.GuildID, reaction.UserID)
			checkErr("Error while converting userID into Member", err)
			for _, RoleID := range member.Roles{
				if RoleID == RoleNeededToAccept{
					switch reaction.Emoji.Name{
						case "❌":
							sendMessage(session, reaction.ChannelID, "Declined "+ Thxs[i].mentioned.Username+ "'s request!")
							Thxs[i] = Thx{Thxs[i].messageID, Declined, Thxs[i].mentioner, Thxs[i].mentioned}
						case "✅":
							sendMessage(session, reaction.ChannelID, "Added " + Thxs[i].mentioned.Username+ " to giveaway!")
							Thxs[i] = Thx{Thxs[i].messageID, Accepted, Thxs[i].mentioner, Thxs[i].mentioned}
							sendMessage(session, reaction.ChannelID, "Participants: " + strings.Join(getAllUsernames(getAllConfirmedUsers(Thxs)), ", "))
					}
					return

				}else{
					if reaction.UserID == session.State.User.ID{
						return
					}
					//hasn't perm
					err := session.MessageReactionRemove(reaction.ChannelID, reaction.MessageID, reaction.Emoji.Name, reaction.UserID)
					checkErr("Error while deleting reaction", err)
				}
			}
		}
	}
}
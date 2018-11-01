package bot

import (
	"github.com/bwmarrin/discordgo"
)

func LuChannel(session *discordgo.Session, channelID string) (channel *discordgo.Channel, err error) {
	channel, err = session.State.Channel(channelID)
	if err == nil {
		return
	}
	channel, err = session.Channel(channelID)
	return
}

func LuMember(session *discordgo.Session, guildID, userID string) (member *discordgo.Member, err error) {
	member, err = session.State.Member(guildID, userID)
	if err == nil {
		return
	}
	member, err = session.GuildMember(guildID, userID)
	return
}

func LuUser(session *discordgo.Session, guildID, userID string) (user *discordgo.User, err error) {
	var member *discordgo.Member
	member, err = session.State.Member(guildID, userID)
	if err == nil {
		user = member.User
		return
	}
	user, err = session.User(userID)
	return
}

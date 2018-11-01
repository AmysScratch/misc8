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

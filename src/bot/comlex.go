package bot

import (
	"github.com/bwmarrin/discordgo"
)

const (
	Com_SUDO = uint64(1) << iota,
	Com_,
)

type Com struct {
	Nick     string
	Username string
	BestName string
	Channel  *discordgo.Channel
	Member   *discordgo.Member
	User     *discordgo.Member
	UID64    int64
	GID64    int64
	Flags    uint64
}

func NewCom(s *discordgo.Session, message *discordgo.Message, payl string) (com *Com, err error) {
	com = new(Com)
	err = com.Init(s, message, payl)
	return
}

func (com *Com) Init(s *discordgo.Session, message *discordgo.Message, payl string) error {
	var err error
	if message != nil {
		if com.Channel, err = LuChannel(s, message.ChannelID); err != nil {
			return
		}
		if com.GID64, err = strconv.ParseInt(com.Channel.GuildID, 10, 64); err != nil {
			com.GID64 = 0
		}
		com.Member, err = LuMember(s, com.Channel.GuildID, message.Author.ID)
		if err == nil {
			com.User = com.Member.User
			com.Nick = com.Member.Nick
			com.BestName = com.Member.Nick
		} else {
			com.Member = nil
			com.User = s.User(message.Author.ID)
		}
		com.Username = com.User.Username
		if len(com.BestName) < 1 {
			com.BestName = com.User.Username
		}
	}
}

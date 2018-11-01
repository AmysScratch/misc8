package bot

import (
	"github.com/bwmarrin/discordgo"
)

const (
	Com_SUDO = uint64(1) << iota,
	Com_,
)

type Com struct {
	Channel *discordgo.Channel
	Member  *discordgo.Member
	User    *discordgo.Member
	UID64   int64
	GID64   int64
	Flags   uint64
}

func NewCom(s *discordgo.Session, message *discordgo.Message, payl string) (com *Com, err error) {
	com = new(Com)
	err = com.Init(s, message, payl)
	return
}

func (com *Com) Init(s *discordgo.Session, message *discordgo.Message, payl string) error {
}

package bot

import (
	"math/rand"
	"time"

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
	Tail     string
	Session  *discordgo.Session
	Message  *discordgo.Message
	Channel  *discordgo.Channel
	Member   *discordgo.Member
	User     *discordgo.User
	UID64    int64
	GID64    int64
	Entropy  uint64
	Flags    uint64
}

func NewCom(s *discordgo.Session, message *discordgo.Message, payl string) (com *Com, err error) {
	com = new(Com)
	err = com.Init(s, message, payl)
	return
}

func (com *Com) Init(s *discordgo.Session, message *discordgo.Message, payl string) error {
	var err error
	com.Session = s
	com.Message = message
	com.Tail = payl
	if com.Channel, err = LuChannel(s, message.ChannelID); err != nil {
		return
	}
	if com.UID64, err = strconv.ParseInt(message.Author.ID, 10, 64); err != nil {
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
		if com.User, err = s.User(message.Author.ID); err != nil {
			return
		}
	}
	com.Username = com.User.Username
	if len(com.BestName) < 1 {
		com.BestName = com.User.Username
	}
	com.Entropy = PRG.Uint64()
	return nil
}

var PRG *rand.Rand

func init() {
	PRG = rand.New(rand.NewSource(time.Now().UnixNano()))
}

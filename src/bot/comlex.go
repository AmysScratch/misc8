package bot

const (
	Com_SUDO = uint64(1) << iota,
	Com_,
)

type Com struct {
	Flags uint64
}

func NewCom(s *discordgo.Session, message *discordgo.Message, payl string) (com *Com, err error) {
	com = new(Com)
	err = com.Init(s, message, payl)
	return
}

func (com *Com) Init(s *discordgo.Session, message *discordgo.Message, payl string) error {
}

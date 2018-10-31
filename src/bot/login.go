package bot

import (
	"database/sql"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

struct Bot {
	Session *discordgo.Session
}

var Bots []*Bot
var DB *sql.DB

func Main() {
	bcfgs, err := ioutil.ReadFile("/usr/local/Tokens/Sa" + "sa8.dat")
	if err != nil {
		panic(err)
	}
	if DB, err := sql.Open("sqlite3", "/usr/local/Sa" + "sa8/Assets/Sa" + "sa8.db"); err != nil {
		panic(err)
	}
	for cfg := range strings.Split(string(bcfgs), "\n") {
		Bots = append(Bots, login(cfg))
	}
}

func login(line string) *Bot {
	bot := new(Bot)
	for kvws := range strings.Split(line, ";") {
		kv := strings.Split(strings.TrimSpace(kvws), "=")
		if len(kv) != 2 {
			return nil
		}
		switch strings.ToLower(kv[0]) {
		case "token":
			session, err := discordgo.New("Bot " + kv[1])
			if err != nil {
				panic("Failed to log in as Bot " + kv[1] + ".")
			}
			bot.Session = session
		}
	}
	if bot.Session != nil {
		return bot
	}
	return nil
}

package bot

import (
	"database/sql"
	"io/ioutil"
	"strings"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

struct Bot {
	Session  *discordgo.Session
	User     *discordgo.User
	OwnerIDs []string
}

var Bots []*Bot
var DB *sql.DB
var AmyIDs = []string{
	"413206608" + "839966721",
}
var Amys []*discordgo.User

func Main() {
	bcfgs, err := ioutil.ReadFile("/usr/local/Tokens/Sa" + "sa8.dat")
	if err != nil {
		panic(err)
	}
	if DB, err := sql.Open("sqlite3", "/usr/local/Sa" + "sa8/Assets/Sa" + "sa8.db"); err != nil {
		panic(err)
	}
	for _, cfg := range strings.Split(string(bcfgs), "\n") {
		Bots = append(Bots, login(cfg))
	}
	if len(Bots) < 1 {
		panic("No bots")
	}
	for _, amyID := range AmyIDs {
		amy, err := Bots[0].Session.User(amyID)
		if err != nil {
			panic("Could not resolve Amy")
		}
		Amys = append(Amys, amy)
	}
	LexerMain()
}

func login(line string) *Bot {
	bot := new(Bot)
	for _, kvws := range strings.Split(line, ";") {
		kv := strings.Split(strings.TrimSpace(kvws), "=")
		if len(kv) != 2 {
			return nil
		}
		switch strings.ToLower(kv[0]) {
		case "owner":
			bot.Owners = append(bot.OwnerIDs, kv[1])
		case "token":
			session, err := discordgo.New("Bot " + kv[1])
			if err != nil {
				panic("Failed to log in as Bot " + kv[1] + ".")
			}
			bot.Session = session
		}
	}
	if bot.Session != nil {
		user, err := bot.Session.User("@me")
		if err != nil {
			panic("Could not resolve @me")
		}
		bot.User = user
		return bot
	}
	return nil
}

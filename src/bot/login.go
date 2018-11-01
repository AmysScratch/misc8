package bot

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall" // TODO: deprecated, use the new unix thingo

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

type Bot struct {
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
	db, err := sql.Open("sqlite3", "/usr/local/Sa" + "sa8/Assets/Sa" + "sa8.db")
	if err != nil {
		panic(err)
	}
	DB = db
	for _, cfg := range strings.Split(string(bcfgs), "\n") {
		if len(cfg) > 0 {
			Bots = append(Bots, login(cfg))
		}
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
	for _, bot := range Bots {
		bot.Session.AddHandler(onMessageCreate)
		bot.Session.AddHandler(onMessageDelete)
		bot.Session.AddHandler(onMessageReactionAdd)
		bot.Session.AddHandler(onMessageReactionRemove)
		bot.Session.AddHandler(onMessageUpdate)
		bot.Session.AddHandler(onGuildMemberAdd)
		bot.Session.AddHandler(onGuildMemberRemove)
		bot.Session.AddHandler(onGuildMemberUpdate)
		bot.Session.AddHandler(onVoiceStateUpdate)
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	for _, bot := range Bots {
		if err := bot.Session.Open(); err != nil {
			panic(err)
		}
		defer bot.Session.Close()
	}
	go cyclePlayingStatus()
	<-sc
}

func login(line string) *Bot {
	success := false
	fmt.Println("===")
	defer func() {
		if success {
			fmt.Println("...")
		} else {
			fmt.Println("... // failed")
		}
	}()
	bot := new(Bot)
	for _, kvws := range strings.Split(line, ";") {
		kvtrimmed := strings.TrimSpace(kvws)
		if len(kvtrimmed) < 1 {
			continue
		}
		kv := strings.Split(kvtrimmed, "=")
		if len(kv) != 2 {
			panic("len(kv) != 2: ``" + kvws + "\u00b4\u00b4")
		}
		understood := true
		isToken := false
		switch strings.ToLower(kv[0]) {
		case "owner":
			bot.OwnerIDs = append(bot.OwnerIDs, kv[1])
		case "token":
			isToken = true
			session, err := discordgo.New("Bot " + kv[1])
			if err != nil {
				panic("Failed to log in as Bot " + kv[1] + ".")
			}
			bot.Session = session
		default:
			understood = false
		}
		if understood {
			showVal := "********"
			if !isToken {
				showVal = kv[1]
			}
			fmt.Println(kv[0] + "\t\t``" + showVal + "\u00b4\u00b4")
		} else {
			fmt.Println("Unknown key ``" + kv[0] + "\u00b4\u00b4")
		}
	}
	if bot.Session == nil {
		panic("bot.Session == nil")
	}
	user, err := bot.Session.User("@me")
	if err != nil {
		panic("Could not resolve @me")
	}
	bot.User = user
	success = true
	return bot
}

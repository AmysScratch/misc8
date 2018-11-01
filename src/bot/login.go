package bot

import (
	"database/sql"
	"io/ioutil"
	"os"
	"signal"
	"strconv"
	"strings"
	"syscall" // TODO: deprecated, use the new unix thingo
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/mattn/go-sqlite3"
)

var Playing = []string{
	"Stay Alive Today"
}

var PlayingMonth = [12][]string{
	{
		"it's January",
	},
	{
		"Happy ForeverAlone's Day",
	},
	{
		"it's March",
	},
	{
		"AoT S3p2 coming out",
	},
	{
		"Don't Give Up!",
	},
	{
		"Happy Pride!",
	},
	{
		"it's July",
	},
	{
		"it's August",
	},
	{
		"it's September",
	},
	{
		"Happy Halloween!",
		"Vote Democrat",
	},
	{
		"Vote Democrat",
	},
	{
		"Happy Holidays!",
	},
}

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
	go cyclePlayingStatus()
	<-sc
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

func cyclePlayingStatus() {
	var servers string
	for {
		servers = strconv.Format(int64(len(session.State.Guilds)), 10) + " Servers"
		now := time.Now()
		month := int(now.Month()) - 1

		for _, bot := range Bots {
			bot.Session.UpdateStatus(0, servers)
		}
		time.Sleep(10 * time.Second)

		for _, playing := range PlayingMonth[month] {
			for _, bot := range Bots {
				bot.Session.UpdateStatus(0, playing)
			}
			time.Sleep(10 * time.Second)
		}

		for _, playing := range Playing {
			for _, bot := range Bots {
				bot.Session.UpdateStatus(0, playing)
			}
			time.Sleep(10 * time.Second)
		}

		for _, playing := range PlayingMonth[month] {
			for _, bot := range Bots {
				bot.Session.UpdateStatus(0, playing)
			}
			time.Sleep(10 * time.Second)
		}

	}
}

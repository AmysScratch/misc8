package bot

import (
	"github.com/bwmarrin/discordgo"
)

var PlayingFrequent = []string{
	"Prefix is , (comma)",
	"Try ,help",
}

var PlayingInfrequent = []string{
	"Stay alive today",
}

var PlayingMonth = [12][]string{
	{
		"It's January",
	},
	{
		"Happy ForeverAlone's Day",
	},
	{
		"It's March",
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
		"It's July",
	},
	{
		"It's August",
	},
	{
		"It's September",
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

func cyclePlayingStatus() {
	var servers string
	var guilds int64
	for {
		guilds = 0
		for _, bot := range Bots {
			guilds += int64(len(bot.Session.State.Guilds))
		}
		servers = strconv.FormatInt(guilds, 10) + " Servers"
		now := time.Now()
		month := int(now.Month()) - 1

		for _, bot := range Bots {
			bot.Session.UpdateStatus(0, servers)
		}
		time.Sleep(10 * time.Second)

		for _, playing := range PlayingFrequent {
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

		for _, playing := range PlayingInfrequent {
			for _, bot := range Bots {
				bot.Session.UpdateStatus(0, playing)
			}
			time.Sleep(10 * time.Second)
		}

		for _, playing := range PlayingFrequent {
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

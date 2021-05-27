package bot

import (
	"errors"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type BotDiscord struct {
	mu sync.Mutex
	C  chan error

	Token   string
	Session *discordgo.Session
	ID      string

	Running bool
}

var DiscordBot BotDiscord = BotDiscord{
	Running: false,
}

func (bot *BotDiscord) Start() {
	bot.mu.Lock()
	defer bot.mu.Unlock()

	if bot.Token == "" {
		bot.C <- errors.New("bot token is undefined")
		return
	}

	if session, err := discordgo.New("Bot " + bot.Token); err != nil {
		bot.C <- errors.New("error creating bot session: " + err.Error())
		return
	} else {
		bot.Session = session
	}

	if u, err := bot.Session.User("@me"); err != nil {
		bot.C <- errors.New("error gettiong bot ID: " + err.Error())
		return
	} else if u.ID == "" {
		bot.C <- errors.New("bot ID is undefined")
		return
	} else {
		bot.ID = u.ID
	}

	LgrDiscord.Println("Adding handlers...")
	bot.Session.AddHandler(messageHandler)

	LgrDiscord.Println("Connecting...")
	if err := bot.Session.Open(); err != nil {
		bot.C <- errors.New("Error opening connection: " + err.Error())
		return
	}

	bot.Running = true

	LgrDiscord.Println("Running..")
	LgrDiscord.Printf("Use ? to run commands\n")

	bot.C <- nil
}

func (bot *BotDiscord) Stop() {
	bot.Running = false
	bot.C <- bot.Session.Close()
}

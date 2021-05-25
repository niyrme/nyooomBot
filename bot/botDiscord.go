package bot

import (
	"errors"

	"github.com/bwmarrin/discordgo"
)

type BotDiscord struct {
	Token   string
	Session *discordgo.Session
	ID      string

	Running bool
}

var DiscordBot BotDiscord = BotDiscord{
	Running: false,
}

func (b *BotDiscord) Start() error {
	// Check Bot
	if b.Token == "" {
		return errors.New("bot token is undefined")
	}

	if session, err := discordgo.New("Bot " + b.Token); err != nil {
		return err
	} else {
		b.Session = session
	}

	if u, err := b.Session.User("@me"); err != nil {
		return err
	} else {
		b.ID = u.ID
	}

	if b.ID == "" {
		return errors.New("bot ID is undefined")
	}

	LgrDiscord.Println("Adding handlers...")
	b.Session.AddHandler(messageHandler)

	LgrDiscord.Println("Connecting...")
	if err := b.Session.Open(); err != nil {
		return errors.New("Error opening connection: %s" + err.Error())
	}

	LgrDiscord.Println("Running..")
	LgrDiscord.Printf("Use ? to run commands\n")

	b.Running = true
	return nil
}

func (b *BotDiscord) Stop() error {
	return b.Session.Close()
}

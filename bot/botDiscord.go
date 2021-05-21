package bot

import (
	"errors"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Prefix  string
	Token   string
	Session *discordgo.Session
	ID      string
}

var BotDiscord Bot = Bot{}

func (b *Bot) Start() error {
	// Check Bot
	if b.Token == "" {
		return errors.New("Bot token is undefined")
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
		return errors.New("Bot ID is undefined")
	}

	log.Println("[INFO] Adding handlers...")
	b.Session.AddHandler(messageHandler)

	log.Println("[INFO] Connecting...")
	if err := b.Session.Open(); err != nil {
		return errors.New("Error opening connection: %s" + err.Error())
	}

	log.Println("[INFO] Running..")
	log.Printf("[INFO] Use %s to run commands\n", b.Prefix)

	return nil
}

func (b *Bot) Stop() error {
	return b.Session.Close()
}

package main

import (
	"errors"
	"sync"

	"github.com/bwmarrin/discordgo"
)

type bot struct {
	mu sync.Mutex

	Token   string
	Session *discordgo.Session
	ID      string
}

func (b *bot) start() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.Token == "" {
		chanErr <- errors.New("bot token is undefined")
	}

	if session, err := discordgo.New("Bot " + b.Token); err != nil {
		chanErr <- errors.New("error creating bot session: " + err.Error())
	} else {
		b.Session = session
	}

	if u, err := b.Session.User("@me"); err != nil {
		chanErr <- errors.New("error gettiong b ID: " + err.Error())
	} else if u.ID == "" {
		chanErr <- errors.New("bot ID is undefined")
	} else {
		b.ID = u.ID
	}

	chanLog <- ("Adding handlers...")
	b.Session.AddHandler(messageCreate)

	chanLog <- ("Connecting...")
	if err := b.Session.Open(); err != nil {
		chanErr <- errors.New("Error opening connection: " + err.Error())
	}
}

func (b *bot) stop() {
	chanErr <- b.Session.Close()
}

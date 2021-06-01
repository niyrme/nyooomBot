package main

import (
	"nyooomBot-Discord/modules"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == Bot.ID {
		return
	}

	p := 16 - len(msg.Author.Username)
	if p <= 1 {
		p = 1
	}
	chanLog <- (strings.Repeat(" ", p) +
		" <" + msg.Author.Username + "> " +
		strings.ReplaceAll(msg.Content, "\n", " \\n "))

	if msg.Content[0] != '?' {
		return
	}

	_cmarg := strings.Split(strings.ToLower(msg.Content[1:]), " ")

	var (
		cmd  string   = _cmarg[0]
		args []string = _cmarg[1:]
	)

	s.ChannelMessageSend(
		msg.ChannelID,
		modules.AnswerCommand(cmd, args),
	)
}

package main

import (
	"nyooomBot-Discord/modules"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var regexCmd *regexp.Regexp = regexp.MustCompile(`^\?(\w+)\s?(\w+)?`)

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

	if matchCmd := regexCmd.FindStringSubmatch(msg.Content); matchCmd != nil {
		var (
			cmd  string   = strings.TrimSpace(matchCmd[1])
			args []string = strings.Split(strings.TrimSpace(matchCmd[2]), " ")
		)

		s.ChannelMessageSend(
			msg.ChannelID,
			modules.AnswerCommand(cmd, args),
		)
	}
}

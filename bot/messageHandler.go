package bot

import (
	"strings"

	mod "nyooomBot/bot/modules"

	"github.com/bwmarrin/discordgo"
)

func messageHandler(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == DiscordBot.ID {
		return
	}

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
		mod.AnswerCommand(cmd, args),
	)

	LgrDiscord.Printf("New command: {\"%s\" - %s}", msg.Content, msg.Author.Username)
}

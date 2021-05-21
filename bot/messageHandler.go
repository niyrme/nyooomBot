package bot

import (
	"fmt"
	"log"
	"nyooomBot/bot/modules"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var commands map[string]modules.Module = map[string]modules.Module{
	"ping": modules.ModPing,
	"help": modules.ModHelp,
	"dice": modules.ModDice,
}

func messageHandler(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == BotDiscord.ID {
		return
	}

	if string(msg.Content[0]) != BotDiscord.Prefix {
		return
	}

	_cmarg := strings.Split(strings.ToLower(msg.Content[1:]), " ")

	var (
		resp string   = ""
		cmd  string   = _cmarg[0]
		args []string = _cmarg[1:]
	)

	if cmd == "how" {
		if len(args) == 0 {
			resp = "Not enough arguments given! Expected: 1; Got: 0"
		} else {
			if args[0] == "how" {
				resp = "`?how {command}`\n`{command}` should NOT include the `?`"
			} else if args[0] == "desc" {
				resp = "`?desc {command}`\n`{command}` should NOT include the `?`"
			} else if _, ok := commands[args[0]]; ok {
				resp = commands[args[0]].How
			} else {
				resp = fmt.Sprintf("Unknown command `%s`", args[0])
			}
		}
	} else if cmd == "desc" {
		if len(args) == 0 {
			resp = "Not enough arguments given! Expected: 1; Got: 0"
		} else {
			if args[0] == "how" {
				resp = "Shows how to use a command"
			} else if args[0] == "desc" {
				resp = "Shows the description of a command"
			} else if _, ok := commands[args[0]]; ok {
				resp = commands[args[0]].Description
			} else {
				resp = fmt.Sprintf("Unknown command `%s`", args[0])
			}
		}
	} else if _, ok := commands[cmd]; ok {
		resp = commands[cmd].Run(msg.Message, args)
	} else {
		resp = fmt.Sprintf("Unknown command `%s`", args[0])
	}

	s.ChannelMessageSend(msg.ChannelID, resp)

	log.Printf("[INFO] New command: {\"%s\" - %s}", msg.Content, msg.Author.Username)
}

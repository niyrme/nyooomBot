package bot

import (
	"fmt"
	"log"
	"strings"

	mod "nyooomBot/bot/modules"
	discord "nyooomBot/bot/modules/Discord"

	"github.com/bwmarrin/discordgo"
)

var commandsDiscord map[string]mod.Module = map[string]mod.Module{
	"ping": discord.ModPing,
	"help": discord.ModHelp,
	"dice": discord.ModDice,
}

func messageHandler(s *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Author.ID == DiscordBot.ID {
		return
	}

	if msg.Content[0] != '/' {
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
				resp = "`/how {command}`\n`{command}` should NOT include the `/`"
			} else if args[0] == "desc" {
				resp = "`/desc {command}`\n`{command}` should NOT include the `/`"
			} else if _, ok := commandsDiscord[args[0]]; ok {
				resp = commandsDiscord[args[0]].How
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
			} else if _, ok := commandsDiscord[args[0]]; ok {
				resp = commandsDiscord[args[0]].Description
			} else {
				resp = fmt.Sprintf("Unknown command `%s`", args[0])
			}
		}
	} else if _, ok := commandsDiscord[cmd]; ok {
		resp = commandsDiscord[cmd].Run(args)
	} else {
		resp = fmt.Sprintf("Unknown command `%s`", args[0])
	}

	s.ChannelMessageSend(msg.ChannelID, resp)

	log.Printf("[INFO] New command: {\"%s\" - %s}", msg.Content, msg.Author.Username)
}

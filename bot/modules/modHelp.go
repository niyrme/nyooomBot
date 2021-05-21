package modules

import "github.com/bwmarrin/discordgo"

var ModHelp Module = Module{
	Description: "",
	How:         "`?help`",
	Run: func(*discordgo.Message, []string) string {
		return "A helpful message!"
	},
}

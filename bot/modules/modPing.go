package modules

import "github.com/bwmarrin/discordgo"

var ModPing Module = Module{
	Description: "Responds with `Pong!`",
	How:         "`?ping`",
	Run: func(*discordgo.Message, []string) string {
		return "`Pong!`"
	},
}

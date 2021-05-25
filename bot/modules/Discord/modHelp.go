package modDiscord

import (
	mod "nyooomBot/bot/modules"
)

var ModHelp mod.Module = mod.Module{
	Description: "",
	How:         "`/help`",
	Run: func(_ []string) string {
		return "A helpful message!"
	},
}

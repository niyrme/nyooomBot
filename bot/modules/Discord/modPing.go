package modDiscord

import (
	mod "nyooomBot/bot/modules"
)

var ModPing mod.Module = mod.Module{
	Description: "Responds with `Pong!`",
	How:         "`/ping`",
	Run: func(_ []string) string {
		return "`Pong!`"
	},
}

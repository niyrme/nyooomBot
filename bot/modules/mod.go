package modules

import "github.com/bwmarrin/discordgo"

type Module struct {
	Description string
	How         string
	Run         func(*discordgo.Message, []string) string
}

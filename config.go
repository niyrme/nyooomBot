package main

import (
	"github.com/pelletier/go-toml"
)

type BotConfigNames struct {
	Discord string
	Twitch  string
}

type BotConfig struct {
	Name          string
	TwitchChannel string
	Names         BotConfigNames
}

type Config struct {
	Bot BotConfig
}

func getConfig() error {
	if tree, err := toml.LoadFile("./config.toml"); err != nil {
		return err
	} else {
		tree.Unmarshal(&cfg)
		return nil
	}
}

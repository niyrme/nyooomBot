package main

import (
	"github.com/pelletier/go-toml"
)

type TokenConfig struct {
	Discord string
}

type Config struct {
	Token TokenConfig
}

func getConfig() error {
	if tree, err := toml.LoadFile("./config.toml"); err != nil {
		return err
	} else {
		tree.Unmarshal(&cfg)
		return nil
	}
}

package modules

import (
	"fmt"
	"strings"
)

var commands []_Module = []_Module{
	&ModCommands,
	&ModDesc,
	&ModDice,
	&ModHelp,
	&ModHow,
	&ModPing,
}

type _Module interface {
	Run(arg []string) (resp string)
	Super() Module
}

type Module struct {
	Keys []string

	Description string
	Name        string
	How         string
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func AnswerCommand(cmd string, args []string) (resp string) {
	resp = ""
	cmd = strings.TrimSpace(cmd)

	for _, command := range commands {
		if contains(command.Super().Keys, cmd) {
			resp = command.Run(args)
		}
	}

	if resp == "" {
		resp = fmt.Sprintf("Unknown command `%s`", cmd)
	}

	return
}

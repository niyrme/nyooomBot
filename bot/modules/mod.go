package modules

import (
	"fmt"
)

var commands []Module = []Module{
	ModDice,
	ModHelp,
	ModPing,
}

type Module struct {
	Keys []string

	Description string
	How         string

	Run func(args []string) (resp string)
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
	if len(args) == 0 {
		args = []string{""}
	}

	for _, c := range commands {
		if cmd == "how" {
			switch args[0] {
			case "":
				resp = "`?how {command}`\n`{command}` should NOT include the `?`"
			case "how":
				resp = "`?how {command}`\n`{command}` should NOT include the `?`"
			case "desc":
				resp = "`?desc {command}`\n`{command}` should NOT include the `?`"
			default:
				if contains(c.Keys, args[0]) {
					resp = c.How
				}
			}
		} else if cmd == "desc" {
			switch args[0] {
			case "":
				resp = "`?desc {command}`\n`{command}` should NOT include the `?`"
			case "how":
				resp = "Shows how to use a command"
			case "desc":
				resp = "Shows the description of a command"
			default:
				if contains(c.Keys, args[0]) {
					resp = c.Description
				}
			}
		} else if contains(c.Keys, cmd) {
			resp = c.Run(args)
		}
	}

	if resp == "" {
		resp = fmt.Sprintf("Unknown command `%s`", cmd)
	}

	return
}

package modules

import (
	"fmt"
	"log"
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

	log.Printf("[DEBUG]   > cmd: %s | args: %v", cmd, args)

	for _, c := range commands {
		if cmd == "how" {
			if args[0] == "" || args[0] == "how" {
				resp = "`?how {command}`\n`{command}` should NOT include the `?`"
			} else if args[0] == "desc" {
				resp = "`?desc {command}`\n`{command}` should NOT include the `?`"
			} else {
				if contains(c.Keys, args[0]) {
					resp = c.How
				} else {
					resp = fmt.Sprintf("Unknown command `%s`", args[0])
				}
			}
		} else if cmd == "desc" {
			if args[0] == "" {
				resp = "`?desc {command}`\n`{command}` should NOT include the `?`"
			} else if args[0] == "how" {
				resp = "Shows how to use a command"
			} else if args[0] == "desc" {
				resp = "Shows the description of a command"
			} else {
				if contains(c.Keys, args[0]) {
					resp = c.Description
				} else {
					resp = fmt.Sprintf("Unknown command `%s`", args[0])
				}
			}
		}

		if contains(c.Keys, cmd) {
			resp = c.Run(args)
		}
	}

	if resp == "" {
		resp = fmt.Sprintf("Unknown command `%s`", cmd)
	}

	return
}

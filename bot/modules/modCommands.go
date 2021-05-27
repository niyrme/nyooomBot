package modules

import "fmt"

var ModCommands Module = Module{
	Keys: []string{
		"c",
		"commands",
	},

	Name:        "Commands",
	Description: "A list of all commands",
	How:         "`?commands` or `?c`",

	Run: func(args []string) (resp string) {
		resp = "?{Commands, How, Desc"
		for _, cmd := range commands {
			resp += fmt.Sprintf(", %s", cmd.Name)
		}
		resp += "}"
		return
	},
}

package modules

import "fmt"

var ModCommands ModuleCommands = ModuleCommands{
	Module{
		Keys: []string{
			"c",
			"commands",
		},

		Name:        "Commands",
		Description: "A list of all commands",
		How:         "`?commands` or `?c`",
	},
}

type ModuleCommands struct {
	Module
}

func (mod *ModuleCommands) Run(args []string) (resp string) {
	resp = "?{"

	for _, cmd := range commands {
		resp += fmt.Sprintf("%s, ", cmd.Super().Name)
	}
	resp = resp[:len(resp)-2] + "}"
	return
}

func (mod *ModuleCommands) Super() Module {
	return mod.Module
}

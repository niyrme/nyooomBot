package modules

var ModDesc ModuleDesc = ModuleDesc{
	Module{
		Keys: []string{
			"desc",
			"description",
		},

		Name:        "Description",
		Description: "Shows the description of a command",
		How:         "`?desc {command}`\n`{command}` should NOT include the `?`",
	},
}

type ModuleDesc struct {
	Module
}

func (mod *ModuleDesc) Run(args []string) (resp string) {
	resp = ""

	if len(args) == 0 || args[0] == "" {
		resp = mod.How
		return
	}

	for _, cmd := range commands {
		if Contains(cmd.Super().Keys, args[0]) {
			resp = cmd.Super().Description
		}
	}

	if resp == "" {
		resp = "Unknown command " + args[0]
	}

	return
}

func (mod *ModuleDesc) Super() Module {
	return mod.Module
}

package modules

var ModHelp ModuleHelp = ModuleHelp{
	Module{
		Keys: []string{
			"h",
			"help",
		},

		Name:        "Help",
		Description: "Shows how to use a command",
		How:         "`?help {command}` or `?h {command}`",
	},
}

type ModuleHelp struct {
	Module
}

func (mod *ModuleHelp) Run(args []string) (resp string) {
	resp = ""

	if len(args) == 0 || args[0] == "" {
		resp = mod.How
		return
	}

	for _, cmd := range commands {
		if Contains(cmd.Super().Keys, args[0]) {
			resp = cmd.Super().How
		}
	}

	if resp == "" {
		resp = "Unknown command " + args[0]
	}

	return
}

func (mod *ModuleHelp) Super() Module {
	return mod.Module
}

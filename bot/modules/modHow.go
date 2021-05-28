package modules

var ModHow ModuleHow = ModuleHow{
	Module{
		Keys: []string{
			"how",
		},

		Name:        "How",
		Description: "Shows how to use a command",
		How:         "`?how {command}`\n`{command}` should NOT include the `?`",
	},
}

type ModuleHow struct {
	Module
}

func (mod *ModuleHow) Run(args []string) (resp string) {
	resp = ""

	if len(args) == 0 || args[0] == "" {
		resp = mod.How
		return
	}

	for _, cmd := range commands {
		if contains(cmd.Super().Keys, args[0]) {
			resp = cmd.Super().How
		}
	}

	if resp == "" {
		resp = "Unknown command " + args[0]
	}

	return
}

func (mod *ModuleHow) Super() Module {
	return mod.Module
}

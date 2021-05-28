package modules

var ModHelp ModuleHelp = ModuleHelp{
	Module{
		Keys: []string{
			"h",
			"help",
		},

		Name:        "Help",
		Description: "",
		How:         "`?help` or `?h`",
	},
}

type ModuleHelp struct {
	Module
}

func (mod *ModuleHelp) Run(args []string) (resp string) {
	return "A helpful message!"
}

func (mod *ModuleHelp) Super() Module {
	return mod.Module
}

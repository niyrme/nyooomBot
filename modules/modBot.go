package modules

var ModBot ModuleBot = ModuleBot{
	Module{
		Keys: []string{
			"bot",
		},

		Name:        "Bot",
		Description: "Tells you where to find my source code",
		How:         "`?bot`",
	},
}

type ModuleBot struct {
	Module
}

func (mod *ModuleBot) Run(args []string) (resp string) {
	return "https://github.com/niyrme/nyooomBot"
}

func (mod *ModuleBot) Super() Module {
	return mod.Module
}

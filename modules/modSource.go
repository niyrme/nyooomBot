package modules

var ModSource ModuleSource = ModuleSource{
	Module{
		Keys: []string{
			"src",
			"source",
		},

		Name:        "Source",
		Description: "Tells you where to find my source code",
		How:         "`?source` or `?src`",
	},
}

type ModuleSource struct {
	Module
}

func (mod *ModuleSource) Run(args []string) (resp string) {
	return "https://github.com/niyrme/nyooomBot"
}

func (mod *ModuleSource) Super() Module {
	return mod.Module
}

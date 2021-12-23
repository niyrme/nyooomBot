package modules

var ModNomic ModuleNomic = ModuleNomic{
	Module{
		Keys: []string{
			"nomic",
		},

		Name:        "Nomic",
		Description: "Why I'm not using a microphone.",
		How:         "`?nomic`",
	},
}

type ModuleNomic struct {
	Module
}

func (mod *ModuleNomic) Run(args []string) (resp string) {
	return "Not using a microphone, because I'm an introverted and antisocial fuck."
}

func (mod *ModuleNomic) Super() Module {
	return mod.Module
}

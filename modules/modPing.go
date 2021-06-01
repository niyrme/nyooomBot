package modules

var ModPing ModulePing = ModulePing{
	Module{
		Keys: []string{
			"ping",
		},

		Name:        "Ping",
		Description: "Responds with 'Pong!'",
		How:         "`?ping`",
	},
}

type ModulePing struct {
	Module
}

func (mod *ModulePing) Run(args []string) (resp string) {
	return "Pong!"
}

func (mod *ModulePing) Super() Module {
	return mod.Module
}

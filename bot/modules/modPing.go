package modules

var ModPing Module = Module{
	Keys: []string{
		"ping",
	},

	Name:        "Ping",
	Description: "Responds with 'Pong!'",
	How:         "`?ping`",

	Run: func(args []string) (resp string) {
		return "Pong!"
	},
}

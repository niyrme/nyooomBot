package modules

var ModHelp Module = Module{
	Keys: []string{
		"h",
		"help",
	},

	Name:        "Help",
	Description: "",
	How:         "`?help` or `?h`",

	Run: func(args []string) (resp string) {
		return "A helpful message!"
	},
}

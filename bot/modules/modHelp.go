package modules

var ModHelp Module = Module{
	Keys: []string{
		"h",
		"help",
	},

	Description: "",
	How:         "`?help`",

	Run: func(args []string) (resp string) {
		return "A helpful message!"
	},
}

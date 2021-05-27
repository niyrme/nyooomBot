package modules

import (
	"fmt"
	"math/rand"
	"strconv"
)

var ModDice Module = Module{
	Keys: []string{
		"d",
		"dice",
	},

	Description: "Roll a Dice of any size!",
	How:         "`?d {size}` or `?dice {size}` | `{size}` must be a number greater than 1 | Ranges from `1` to `9223372036854775807 (2^63 - 1)`",

	Run: func(args []string) (resp string) {
		resp = ""

		if len(args) == 0 {
			resp = "Not enough arguments given! Expected: 1; Got: 0"
		} else if args[0] == "" {
			resp = "Not enough arguments given! Expected: 1; Got: 0"
		}

		if resp == "" {
			if roll, err := strconv.ParseInt(args[0], 10, 64); err != nil {
				resp = "Could not parse given argument"
			} else {
				if roll < 1 {
					resp = "Argument `{size}` must be a number geater than 1"
				} else {
					resp = fmt.Sprintf("Rolled `%v`!", rand.Intn(int(roll))+1)
				}
			}
		}

		return
	},
}

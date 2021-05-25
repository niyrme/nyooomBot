package modTwitch

import (
	"fmt"
	"math/rand"
	"strconv"

	mod "nyooomBot/bot/modules"
)

var ModDice mod.Module = mod.Module{
	Run: func(args []string) string {
		if len(args) == 0 {
			return "Not enough arguments given! Expected: 1; Got: 0"
		}

		if roll, err := strconv.ParseInt(args[0], 10, 64); err != nil {
			return "Could not parse given argument"
		} else {
			if roll < 1 {
				return "Argument `{size}` must be a number geater than 1"
			} else {
				return fmt.Sprintf("Rolled a `%v`!", rand.Intn(int(roll))+1)
			}
		}
	},
}

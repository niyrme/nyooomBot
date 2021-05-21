package modules

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

var ModDice Module = Module{
	Description: "Roll a Dice of any size!",
	How:         "`?dice {size}`\n`{size}` must be a number greater than 1\nRanges from `1` to `9223372036854775807 (2^63 - 1)`",
	Run: func(_ *discordgo.Message, args []string) string {
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

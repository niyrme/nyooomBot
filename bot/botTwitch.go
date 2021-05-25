package bot

import (
	"regexp"
	"strings"
	"time"

	mod "nyooomBot/bot/modules"
	modTwitch "nyooomBot/bot/modules/Twitch"

	"github.com/gempir/go-twitch-irc/v2"
)

var (
	/// Regexes
	regexCmd *regexp.Regexp = regexp.MustCompile(`^\?(\w+)\s?(\w+)?`)

	commandsTwitch map[string]mod.Module = map[string]mod.Module{
		"h":    modTwitch.ModH,
		"dice": modTwitch.ModDice,
	}
)

func StartBot(channel, oAuth string) (client *twitch.Client) {
	client = twitch.NewClient("nyooomBot", oAuth)

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		LgrTwitch.Printf("New message by `%s`: `%s`", message.User.Name, message.Message)

		if matchCmd := regexCmd.FindStringSubmatch(message.Message); matchCmd != nil {
			cmd := strings.TrimSpace(matchCmd[1])
			args := strings.Split(strings.TrimSpace(matchCmd[2]), " ")

			LgrTwitch.Printf("New command by `%s`: (%v) | (%v)", message.User.Name, cmd, args)

			if cmd == "disconnect" {
				if message.User.Name == channel {
					client.Say(channel, `'ight imma head out`)
					time.Sleep(3 * time.Second)
					client.Disconnect()
				} else {
					client.Say(channel, "You have nothing to tell me!")
				}
			} else if _, ok := commandsTwitch[cmd]; ok {
				commandsTwitch[cmd].Run(args)
			} else {
				client.Say(channel, "Unknown command: "+cmd)
			}
		}
	})

	return
}

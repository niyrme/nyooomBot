package bot

import (
	"regexp"
	"strings"
	"sync"
	"time"

	mod "nyooomBot/bot/modules"

	"github.com/gempir/go-twitch-irc/v2"
)

var regexCmd *regexp.Regexp = regexp.MustCompile(`^\?(\w+)\s?(\w+)?`)

type BotTwitch struct {
	mu sync.Mutex
	C  chan error

	Token   string
	Client  *twitch.Client
	Channel string

	Running bool
}

var TwitchBot BotTwitch = BotTwitch{
	Running: false,
}

func (bot *BotTwitch) Start() {
	bot.mu.Lock()

	bot.Client = twitch.NewClient("nyooomBot", bot.Token)
	bot.Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		LgrTwitch.Printf(
			"%16s %s",
			"<"+message.User.Name+">",
			strings.ReplaceAll(message.Message, "\n", " \\n "),
		)

		if matchCmd := regexCmd.FindStringSubmatch(message.Message); matchCmd != nil {
			cmd := strings.TrimSpace(matchCmd[1])
			args := strings.Split(strings.TrimSpace(matchCmd[2]), " ")

			if cmd == "disconnect" {
				if message.User.Name == bot.Channel {
					bot.Client.Say(bot.Channel, `'ight imma head out`)
					time.Sleep(3 * time.Second)
					bot.Stop()
					return
				} else {
					bot.Client.Say(bot.Channel, "You have nothing to tell me!")
				}
			}

			bot.Client.Say(
				bot.Channel,
				mod.AnswerCommand(cmd, args),
			)
		}
	})

	bot.Running = true
	bot.mu.Unlock()

	go bot.Client.Connect() // TODO: find a way to handle a potential error
	go bot.Client.Join("niyrme")

	LgrTwitch.Println("Running..")
	LgrTwitch.Printf("Use ? to run commands\n")

	bot.C <- nil
}

func (bot *BotTwitch) Stop() {
	bot.Running = false
	bot.C <- bot.Client.Disconnect()
}

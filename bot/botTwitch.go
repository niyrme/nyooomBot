package bot

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	mod "nyooomBot/bot/modules"
	"nyooomBot/logging"

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

func colorMsg(hex, msg string) string {
	hex = strings.TrimSpace(hex)

	var err error
	strR, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		return msg
	}
	strG, err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil {
		return msg
	}
	strB, err := strconv.ParseInt(hex[4:6], 16, 64)
	if err != nil {
		return msg
	}

	strR = strR * 100 / 255
	strG = strG * 100 / 255
	strB = strB * 100 / 255

	return fmt.Sprintf(
		"\033[38;2;%v;%v;%vm%s\033[0m",
		strR,
		strG,
		strB,
		msg,
	)
}

func (bot *BotTwitch) Start() {
	bot.mu.Lock()

	bot.Client = twitch.NewClient("nyooomBot", bot.Token)
	bot.Client.OnConnect(func() {
		logging.LogTwitch("Connected.")
	})
	bot.Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		var msg string
		if len(message.User.Color) == 7 {
			msg = colorMsg(message.User.Color[1:], "<"+message.User.Name+">")
		} else {
			msg = "<" + message.User.Name + ">"
		}
		p := 16 - len(message.User.Name)
		if p <= 1 {
			p = 1
		}
		logging.LogTwitch(
			strings.Repeat(" ", p) +
				" " + msg + " " + // because ansi colors delete spaces
				message.Message)

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

	logging.LogTwitch("Running..")

	bot.C <- nil
}

func (bot *BotTwitch) Stop() {
	bot.Running = false
	bot.C <- bot.Client.Disconnect()
}

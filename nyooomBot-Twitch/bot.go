package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"nyooomBot-Twitch/modules"

	"github.com/gempir/go-twitch-irc/v2"
)

var regexCmd *regexp.Regexp = regexp.MustCompile(`^\?(\w+)\s?(\w+)?`)

type bot struct {
	mu sync.Mutex

	Token   string
	Client  *twitch.Client
	Channel string
}

func (b *bot) start() {
	b.mu.Lock()

	b.Client = twitch.NewClient("nyooomBot", b.Token)
	b.Client.OnConnect(func() {
		chanLog <- "Connected."
		chanLog <- "Running..."
	})
	b.Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if strings.Contains(message.Message, "bigfollows . com") {
			b.Client.Say(b.Channel, fmt.Sprintf("/ban @%s no", message.User.Name))
			return
		}
		var user string = "<" + message.User.DisplayName + ">"
		if len(message.User.Color) == 7 {
			user = colorMsg(message.User.Color[1:], user)
		}

		padding := 16 - len(message.User.Name)
		if padding <= 1 {
			padding = 1
		}

		chanLog <- (strings.Repeat(" ", padding) +
			" " + user + " " + // because ansi colors delete spaces
			message.Message)

		if matchCmd := regexCmd.FindStringSubmatch(message.Message); matchCmd != nil {
			cmd := strings.TrimSpace(matchCmd[1])
			args := strings.Fields(strings.TrimSpace(matchCmd[2]))

			if cmd == "disconnect" {
				if message.User.Name == b.Channel {
					b.Client.Say(b.Channel, `'ight imma head out`)
					time.Sleep(3 * time.Second)
					b.stop()
					return
				} else {
					b.Client.Say(b.Channel, "You have nothing to tell me!")
				}
			}

			b.Client.Say(
				b.Channel,
				modules.AnswerCommand(cmd, args),
			)
		}
	})

	b.mu.Unlock()

	go b.Client.Connect()
	go b.Client.Join(b.Channel)

	chanErr <- nil
}

func (b *bot) stop() {
	chanErr <- b.Client.Disconnect()
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

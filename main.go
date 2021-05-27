package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	bot "nyooomBot/bot"

	"github.com/joho/godotenv"
)

var cfg Config = Config{}

func main() {
	if err := getConfig(); err != nil {
		log.Fatalf("Error loading config.toml file! %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file! %s", err.Error())
	}

	bot.CreateLogger(bot.LGRs{
		Discord: cfg.Bot.Names.Discord,
		Twitch:  cfg.Bot.Names.Twitch,
	})

	var (
		/// channels
		chanDiscord chan error = make(chan error)
		chanTwitch  chan error = make(chan error)
	)

	bot.DiscordBot.C = chanDiscord
	bot.TwitchBot.C = chanTwitch

	bot.DiscordBot.Token = os.Getenv("DISCORD_TOKEN")
	bot.TwitchBot.Token = os.Getenv("TWITCH_TOKEN")

	bot.TwitchBot.Channel = cfg.Bot.TwitchChannel

	go bot.DiscordBot.Start()
	go bot.TwitchBot.Start()

	if err := <-bot.DiscordBot.C; err != nil {
		bot.LgrDiscord.Printf("Error starting bot! %s", err.Error())
	}

	if err := <-bot.TwitchBot.C; err != nil {
		bot.LgrTwitch.Printf("Error starting bot! %s", err.Error())
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if bot.DiscordBot.Running {
		go bot.DiscordBot.Stop()

		if err := <-bot.DiscordBot.C; err != nil {
			bot.LgrDiscord.Printf("Error stopping bot cleanly! %s", err.Error())
		} else {
			bot.LgrDiscord.Println("Bot stopped successfully")
		}
	}

	if bot.TwitchBot.Running {
		go bot.TwitchBot.Stop()

		if err := <-bot.TwitchBot.C; err != nil {
			bot.LgrTwitch.Printf("Error stopping bot cleanly! %s", err.Error())
		} else {
			bot.LgrTwitch.Println("Bot stopped successfully")
		}
	}
}

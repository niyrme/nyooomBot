package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"nyooomBot/bot"
	"nyooomBot/logging"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file! %s", err.Error())
	}

	var (
		/// channels
		chanDiscord chan error = make(chan error, 128)
		chanTwitch  chan error = make(chan error, 128)
	)

	bot.DiscordBot.C = chanDiscord
	bot.TwitchBot.C = chanTwitch

	bot.DiscordBot.Token = os.Getenv("DISCORD_TOKEN")
	bot.TwitchBot.Token = os.Getenv("TWITCH_TOKEN")

	bot.TwitchBot.Channel = os.Getenv("TWITCH_CHANNEL")

	go bot.DiscordBot.Start()
	go bot.TwitchBot.Start()

	if err := <-bot.DiscordBot.C; err != nil {
		logging.LogDiscord("Error starting bot! " + err.Error())
	}

	if err := <-bot.TwitchBot.C; err != nil {
		logging.LogTwitch("Error starting bot! " + err.Error())
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	logging.Info("Stopping bots...")

	if bot.DiscordBot.Running {
		go bot.DiscordBot.Stop()

		if err := <-bot.DiscordBot.C; err != nil {
			logging.LogDiscord("Error stopping bot cleanly! " + err.Error())
		} else {
			logging.LogDiscord("Bot stopped successfully")
		}
	}

	if bot.TwitchBot.Running {
		go bot.TwitchBot.Stop()

		if err := <-bot.TwitchBot.C; err != nil {
			logging.LogTwitch("Error stopping bot cleanly! " + err.Error())
		} else {
			logging.LogTwitch("Bot stopped successfully")
		}
	}
}

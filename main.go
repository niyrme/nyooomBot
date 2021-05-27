package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	bot "nyooomBot/bot"

	"github.com/joho/godotenv"
)

var (
	cfg Config = Config{}
)

func main() {
	if err := getConfig(); err != nil {
		log.Fatalf("Error loading config.toml file! %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file! %s", err.Error())
	}

	bot.DiscordBot.Token = os.Getenv("DISCORD_TOKEN")
	if err := bot.DiscordBot.Start(); err != nil {
		bot.LgrDiscord.Fatalf("Error starting bot! %s", err.Error())
	} else {
		bot.LgrDiscord.Println("Bot is running")
	}

	client := bot.StartBot(cfg.Bot.TwitchChannel, cfg.Bot.Name, os.Getenv("TWITCH_TOKEN"))
	client.Join("niyrme")
	if err := client.Connect(); err != nil {
		bot.LgrTwitch.Fatalf("Failed to connect to twitch")
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	if bot.DiscordBot.Running {
		bot.LgrDiscord.Println("Stopping bot...")
		if err := bot.DiscordBot.Stop(); err != nil {
			bot.LgrDiscord.Fatalf("Error stopping discord bot cleanly! %s", err.Error())
		}
		bot.LgrDiscord.Println("Bot stopped successfully")
	}

	if err := client.Disconnect(); err != nil {
		bot.LgrTwitch.Fatalf("Error stopping twitch bot cleanly! %s", err.Error())
	}
	bot.LgrTwitch.Println("Bot stopped successfully")
}

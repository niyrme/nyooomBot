package main

import (
	"log"
	"nyooomBot/bot"
	"os"
	"os/signal"
	"syscall"
)

var (
	cfg Config = Config{}

	// Logger
	lgrInfo *log.Logger
	// lgrWarn *log.Logger
	lgrErr *log.Logger
)

func main() {
	lgrInfo = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lmsgprefix)
	// lgrWarn = log.New(os.Stdout, "[WARN]", log.Ldate|log.Ltime|log.Lmsgprefix)
	lgrErr = log.New(os.Stdout, "[ERR] ", log.Ldate|log.Ltime|log.Lmsgprefix)

	if err := getConfig(); err != nil {
		lgrErr.Fatalf("Error loading config file! %s", err.Error())
	}

	bot.BotDiscord.Prefix = "?"
	bot.BotDiscord.Token = cfg.Token.Discord

	if err := bot.BotDiscord.Start(); err != nil {
		lgrErr.Fatalf("Error starting bot! %s", err.Error())
	}

	lgrInfo.Println("Bot is running...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	lgrInfo.Println("Stopping bot...")
	if err := bot.BotDiscord.Stop(); err != nil {
		lgrErr.Fatalf("Error stopping bot cleanly! %s", err.Error())
	}
	lgrInfo.Println("Bot stopped successfully")
}

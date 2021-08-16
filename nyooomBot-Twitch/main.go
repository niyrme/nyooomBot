package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	Bot bot = bot{}
	/// Channels
	chanErr chan error  = make(chan error, 16)
	chanLog chan string = make(chan string, 128)
	chanMod chan string = make(chan string, 32)
)

func main() {
	if v := os.Getenv("TOKEN"); v == "" {
		log.Fatalln("Error loading TOKEN from environment")
	} else {
		Bot.Token = v
	}
	if v := os.Getenv("CHANNEL"); v == "" {
		log.Fatalln("Error loading CHANNEL from environment")
	} else {
		Bot.Channel = v
	}

	// Because you can't do go func() { }
	go listenErr()
	go listenLog()
	go listenMod()

	go Bot.start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	chanLog <- "Stopping bot..."
	Bot.stop()
}

func listenLog() {
	for {
		lgrInfo.Println(<-chanLog)
	}
}

func listenMod() {
	for {
		lgrMod.Println(<-chanMod)
	}
}

func listenErr() {
	for {
		if err := <-chanErr; err != nil {
			log.Fatalf("An error occured: %s", err.Error())
		}
	}
}

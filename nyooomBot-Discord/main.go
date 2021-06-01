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
)

func main() {
	if v := os.Getenv("TOKEN"); v == "" {
		log.Fatalln("Error loading TOKEN from environment")
	} else {
		Bot.Token = v
	}

	go listenErr()
	go listenLog()

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

func listenErr() {
	for {
		if err := <-chanErr; err != nil {
			log.Fatalf("An error occured: %s", err.Error())
		}
	}
}

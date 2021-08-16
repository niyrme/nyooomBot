package main

import (
	"fmt"
	"log"
	"os"
)

var (
	lgrInfo *log.Logger = log.New(os.Stdout, fmt.Sprintf("[%-5s] ", "INFO"), log.Ltime|log.Lmsgprefix)
	lgrErr  *log.Logger = log.New(os.Stdout, fmt.Sprintf("[%-5s] ", "ERR"), log.Ltime|log.Lmsgprefix)
	lgrMod  *log.Logger = log.New(os.Stdout, fmt.Sprintf("[%-5s] ", "MOD"), log.Ltime|log.Lmsgprefix)
)

func LogErr(msg string) {
	lgrErr.Printf(msg)
}

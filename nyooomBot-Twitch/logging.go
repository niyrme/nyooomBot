package main

import (
	"fmt"
	"log"
	"os"
)

var (
	lgrInfo *log.Logger = log.New(os.Stdout, fmt.Sprintf("[%-8s] ", "INFO"), log.Ltime|log.Lmsgprefix)
	lgrErr  *log.Logger = log.New(os.Stdout, fmt.Sprintf("[%-8s] ", "ERR"), log.Ltime|log.Lmsgprefix)
)

func LogErr(msg string) {
	lgrErr.Printf(msg)
}

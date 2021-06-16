package main

import (
	"log"
	"os"
)

var (
	lgrInfo *log.Logger = log.New(os.Stdout, "[INFO] ", log.Ltime|log.Lmsgprefix)
	lgrErr  *log.Logger = log.New(os.Stdout, "[ERR]  ", log.Ltime|log.Lmsgprefix)
)

func LogErr(msg string) {
	lgrErr.Printf(msg)
}

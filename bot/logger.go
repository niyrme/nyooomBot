package bot

import (
	"log"
	"os"
)

var (
	LgrTwitch  *log.Logger = log.New(os.Stdout, "[TWITCH]  > ", log.Ldate|log.Ltime|log.Lmsgprefix)
	LgrDiscord *log.Logger = log.New(os.Stdout, "[DISCROD] > ", log.Ldate|log.Ltime|log.Lmsgprefix)
)

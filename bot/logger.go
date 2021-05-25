package bot

import (
	"log"
	"os"
)

var (
	LgrDiscord *log.Logger = log.New(os.Stdout, "[DISCROD] > ", log.Ldate|log.Ltime|log.Lmsgprefix)
)

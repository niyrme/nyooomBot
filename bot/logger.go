package bot

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	LgrDiscord *log.Logger
	LgrTwitch  *log.Logger
)

type LGRs struct {
	/// Names
	Discord string
	Twitch  string
}

func (l *LGRs) names() []string {
	return []string{
		l.Discord,
		l.Twitch,
	}
}

func CreateLogger(lgr LGRs) {
	max := -1

	for _, name := range lgr.names() {
		if len(name) > max {
			max = len(name)
		}
	}

	LgrDiscord = log.New(
		os.Stdout,
		fmt.Sprintf(
			"%s%s ",
			"["+lgr.Discord+"]",
			strings.Repeat(" ", max-len(lgr.Discord)),
		),
		log.Ldate|log.Ltime|log.Lmsgprefix,
	)

	LgrTwitch = log.New(
		os.Stdout,
		fmt.Sprintf(
			"%s%s ",
			"["+lgr.Twitch+"]",
			strings.Repeat(" ", max-len(lgr.Twitch)),
		),
		log.Ldate|log.Ltime|log.Lmsgprefix,
	)
}

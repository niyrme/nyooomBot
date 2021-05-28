package logging

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	LgrDiscord *log.Logger
	LgrTwitch  *log.Logger
	LgrInfo    *log.Logger
)

type LGRs struct {
	/// Names
	Discord string
	Twitch  string
	Info    string
}

func (l *LGRs) names() []string {
	return []string{
		l.Discord,
		l.Twitch,
		l.Info,
	}
}

var (
	max int  = -1
	Lgr LGRs = LGRs{
		Discord: "DISCORD",
		Twitch:  "TWITCH",
		Info:    "INFO",
	}
)

func Info(msg string) {
	if max == -1 {
		getMax()
	}
	if LgrInfo == nil {
		LgrInfo = log.New(
			os.Stdout,
			fmt.Sprintf(
				"%s%s ",
				"["+Lgr.Info+"]",
				strings.Repeat(" ", max-len(Lgr.Info)),
			),
			log.Ldate|log.Ltime|log.Lmsgprefix,
		)
	}

	LgrInfo.Println(msg)
}

func LogDiscord(msg string) {
	if max == -1 {
		getMax()
	}
	if LgrDiscord == nil {
		LgrDiscord = log.New(
			os.Stdout,
			fmt.Sprintf(
				"%s%s ",
				"["+Lgr.Discord+"]",
				strings.Repeat(" ", max-len(Lgr.Discord)),
			),
			log.Ldate|log.Ltime|log.Lmsgprefix,
		)
	}

	LgrDiscord.Println(msg)
}

func LogTwitch(msg string) {
	if max == -1 {
		getMax()
	}
	if LgrTwitch == nil {
		LgrTwitch = log.New(
			os.Stdout,
			fmt.Sprintf(
				"%s%s ",
				"["+Lgr.Twitch+"]",
				strings.Repeat(" ", max-len(Lgr.Twitch)),
			),
			log.Ldate|log.Ltime|log.Lmsgprefix,
		)
	}

	LgrTwitch.Println(msg)
}

func getMax() {
	max = -1
	for _, name := range Lgr.names() {
		if len(name) > max {
			max = len(name)
		}
	}

	Info("Use ? to run commands")
}

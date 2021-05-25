package modDiscord

type Module struct {
	Description string
	How         string

	Run func([]string) string
}

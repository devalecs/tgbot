package main

import (
	"log"

	"gopkg.in/devalecs/tgbot.v1"
)

func main() {
	b := tgbot.NewBot(tgbot.Config{
		Token: "yourBotAPIToken",
	})

	b.RegisterPlugins(
		new(PingPlugin),
	)

	err := b.Start()
	if err != nil {
		log.Fatal(err)
	}
}

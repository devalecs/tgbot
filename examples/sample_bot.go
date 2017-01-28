package main

import (
	"log"

	"gopkg.in/devalecs/hugo.v1"
)

func main() {
	b := hugo.NewBot(hugo.Config{
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

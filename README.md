# Pluggable Telegram bot written in Golang

[![Travis](https://travis-ci.org/devalecs/tgbot.svg)](https://travis-ci.org/devalecs/tgbot)
[![GoDoc](https://godoc.org/github.com/devalecs/tgbot?status.svg)](https://godoc.org/github.com/devalecs/tgbot)

### Plugin example
```go
package main

import (
	"gopkg.in/devalecs/tgbot.v1"
	"gopkg.in/devalecs/tgo.v1"
)

type PingPlugin struct{}

func (p *PingPlugin) Reply(b tgbot.B, update *tgo.Update) {
	b.SendMessage(tgo.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "pong",
	})
}

func (p *PingPlugin) Match(b tgbot.B, update *tgo.Update) bool {
	return b.MessageTextMatch(update, "ping")
}

func (p *PingPlugin) Settings() tgbot.PluginSettings {
	return tgbot.PluginSettings{
		Help: "*/ping*: Sample ping pong plugin",
	}
}
```
### Bot example
```go
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
```
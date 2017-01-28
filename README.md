# Pluggable Telegram bot written in Golang

[![Travis](https://travis-ci.org/devalecs/hugo.svg)](https://travis-ci.org/devalecs/hugo)
[![GoDoc](https://godoc.org/github.com/devalecs/hugo?status.svg)](https://godoc.org/github.com/devalecs/hugo)

### Plugin example
```go
package main

import (
	"gopkg.in/devalecs/hugo.v1"
	"gopkg.in/devalecs/tgo.v1"
)

type PingPlugin struct{}

func (p *PingPlugin) Reply(b hugo.B, update *tgo.Update) {
	b.SendMessage(tgo.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "pong",
	})
}

func (p *PingPlugin) Match(b hugo.B, update *tgo.Update) bool {
	return b.MessageTextMatch(update, "ping")
}

func (p *PingPlugin) Settings() hugo.PluginSettings {
	return hugo.PluginSettings{
		Help: "*/ping*: Sample ping pong plugin",
	}
}
```
### Bot example
```go
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
```
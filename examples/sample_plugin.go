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

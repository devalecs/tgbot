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

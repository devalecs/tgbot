package tgbot

import (
	"log"

	"gopkg.in/devalecs/tgo.v1"
)

type Bot struct {
	*tgo.Client
	*Util

	config  Config
	plugins []Plugin
}

type Config struct {
	Token         string
	AdminUsername string
}

func NewBot(config Config) *Bot {
	return &Bot{
		config: config,
		Client: tgo.NewClient(config.Token),
	}
}

func (b *Bot) GetPlugins() []Plugin {
	return b.plugins
}

func (b *Bot) UserIsAdmin(user tgo.User) bool {
	return b.config.AdminUsername == user.Username
}

func (b *Bot) RegisterPlugins(plugins ...Plugin) {
	for _, plugin := range plugins {
		b.registerPlugin(plugin)
	}
}

func (b *Bot) Start() error {
	user, err := b.GetMe()
	if err != nil {
		return err
	}

	log.Printf("Starting %s\n", user.Username)

	updatesChan := b.GetUpdatesChan(tgo.GetUpdatesParams{
		Timeout: 60,
	})

	for update := range updatesChan {
		go func() {
			b.runPlugins(update)
		}()
	}

	return nil
}

func (b *Bot) registerPlugin(plugin Plugin) {
	b.plugins = append(b.plugins, plugin)
}

func (b *Bot) runPlugins(update *tgo.Update) {
	for _, p := range b.plugins {
		if p.Match(b, update) {
			p.Reply(b, update)
			break
		}
	}
}

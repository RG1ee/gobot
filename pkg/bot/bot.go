package bot

import (
	"time"

	"log/slog"

	"github.com/spf13/viper"
	tele "gopkg.in/telebot.v3"
)

type TelegramBot struct {
	Bot *tele.Bot
}

func NewTelegramBot() (*TelegramBot, error) {
	bot, err := createBot()
	if err != nil {
		return nil, err
	}

	return &TelegramBot{Bot: bot}, nil
}

func createBot() (*tele.Bot, error) {
	pref := tele.Settings{
		Token:   viper.GetString("TOKEN"),
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
		Verbose: viper.GetBool("DEBUG"),
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

func Start() error {
	tb, err := NewTelegramBot()
	if err != nil {
		return err
	}

	slog.Info("Bot started")
	tb.Bot.Start()
	return nil
}

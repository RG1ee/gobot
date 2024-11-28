package bot

import (
	"time"

	"log/slog"

	ph "github.com/RG1ee/gobot/pkg/bot/handlers/public"
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
		Verbose: viper.GetBool("VERBOSE"),
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	return bot, nil
}

func (tb *TelegramBot) registerHandler() {
	public := tb.Bot.Group()
	public.Handle("/start", ph.HandleStart())
}

func Start() error {
	tb, err := NewTelegramBot()
	if err != nil {
		return err
	}

	tb.registerHandler()

	slog.Info("Bot started")
	tb.Bot.Start()
	return nil
}

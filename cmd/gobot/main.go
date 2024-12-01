package main

import (
	"log/slog"

	"github.com/RG1ee/gobot/internal/bot"
)

func main() {
	slog.Info("Bot started")
	bot.Start()
}

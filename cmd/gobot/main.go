package main

import (
	"log/slog"

	"github.com/RG1ee/gobot/internal/configs"
	"github.com/RG1ee/gobot/pkg/bot"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	err = bot.Start()
	if err != nil {
		slog.Error("Failed to start bot", "Fatal", err)
	}
}

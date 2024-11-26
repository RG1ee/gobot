package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("No .env file found")
    }
}

func main() {
	token, exists := os.LookupEnv("TOKEN")
	if !exists {
		log.Fatal("TOKEN is not set")
	}

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello! It's a bot")
	})

	bot.Start()
}

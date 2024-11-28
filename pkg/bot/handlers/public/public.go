package public

import (
	tele "gopkg.in/telebot.v3"
)

func HandleStart() func(tele.Context) error {
	return func(c tele.Context) error {
		return c.Send("Hello World")
	}
}

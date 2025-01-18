package middleware

import (
	"github.com/RG1ee/gobot/internal/repository"
	tele "gopkg.in/telebot.v3"
)

func Repository(db_backend repository.Cloth) func(next tele.HandlerFunc) tele.HandlerFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			c.Set("repository", db_backend)
			return next(c)
		}
	}
}

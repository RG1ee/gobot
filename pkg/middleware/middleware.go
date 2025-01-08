package middleware

import (
	// stateconst "github.com/RG1ee/gobot/internal/bot/state_const"

	stateconst "github.com/RG1ee/gobot/internal/bot/state_const"
	"github.com/RG1ee/gobot/internal/repository"
	"github.com/avi-gecko/fsm/pkg/fsm"
	tele "gopkg.in/telebot.v3"
)

func FsmMiddleware(fsm fsm.FSM) func(next tele.HandlerFunc) tele.HandlerFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			c.Set("fsm", fsm)
			return next(c)
		}
	}
}

func Repository(db_backend repository.Cloth) func(next tele.HandlerFunc) tele.HandlerFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			c.Set("repository", db_backend)
			return next(c)
		}
	}
}

func CleanupMessages() func(next tele.HandlerFunc) tele.HandlerFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			err := next(c)
			if err != nil {
				return err
			}
			f := c.Get("fsm").(fsm.FSM)
			result, err := f.GetState(uint64(c.Chat().ID))
			if err != nil {
				return nil
			}
			currentState := result.(stateconst.State)
			if currentState.PreviousMessages != nil && len(*currentState.PreviousMessages) > 0 {
				c.Bot().DeleteMany(*currentState.PreviousMessages)
			}
			currentState.PreviousMessages = &currentState.CurrentMessages
			f.SetState(uint64(c.Chat().ID), currentState)
			return nil
		}
	}
}

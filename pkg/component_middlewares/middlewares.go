package component_middlewares

import (
	"github.com/avi-gecko/fsm/pkg/fsm"
	tele "gopkg.in/telebot.v3"
)

func FsmMiddleware(fsm fsm.FSM[State]) func(next tele.HandlerFunc) tele.HandlerFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			c.Set("fsm", fsm)
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
			f := c.Get("fsm").(fsm.FSM[State])
			currentState, err := f.GetState(uint64(c.Chat().ID))
			if err != nil {
				return nil
			}
			if currentState.PreviousMessages != nil && len(*currentState.PreviousMessages) > 0 {
				c.Bot().DeleteMany(*currentState.PreviousMessages)
			}
			currentState.PreviousMessages = &currentState.CurrentMessages
			f.SetState(uint64(c.Chat().ID), currentState)
			return nil
		}
	}
}

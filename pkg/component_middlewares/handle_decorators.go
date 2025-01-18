package component_middlewares

import (
	"github.com/avi-gecko/fsm/pkg/fsm"
	tele "gopkg.in/telebot.v3"
)

func StateGate(c func(tele.Context) error, state UserState) func(tele.Context) error {
	return func(ctx tele.Context) error {
		f := ctx.Get("fsm").(fsm.FSM[State])
		currentState, err := f.GetState(uint64(ctx.Chat().ID))
		if err != nil {
			panic(err)
		}
		if state != currentState.UserState {
			return nil
		}
		return c(ctx)
	}
}

func SaveLastMessage(c func(tele.Context) ([]tele.Editable, error)) func(tele.Context) error {
	return func(ctx tele.Context) error {
		f := ctx.Get("fsm").(fsm.FSM[State])
		ctx.Delete()
		currentState, err := f.GetState(uint64(ctx.Chat().ID))
		if err != nil {
			currentState.UserState = NullState
			currentState.PreviousMessages = nil
			f.SetState(uint64(ctx.Chat().ID), currentState)
		}
		messages, err := c(ctx)
		if err != nil {
			return err
		}
		currentState, _ = f.GetState(uint64(ctx.Chat().ID))
		currentState.CurrentMessages = messages
		f.SetState(uint64(ctx.Chat().ID), currentState)
		return nil
	}
}

package handlerdecorator

import (
	stateconst "github.com/RG1ee/gobot/internal/bot/state_const"
	userstate "github.com/RG1ee/gobot/pkg/user_state"
	"github.com/avi-gecko/fsm/pkg/fsm"
	tele "gopkg.in/telebot.v3"
)

func StateHandler(c func(tele.Context) error, state userstate.UserState) func(tele.Context) error {
	return func(ctx tele.Context) error {
		f := ctx.Get("fsm").(fsm.FSM)
		result, err := f.GetState(uint64(ctx.Chat().ID))
		if err != nil {
			return nil
		}
		currentState, ok := result.(stateconst.State)
		if !ok {
			panic(ok)
		}
		if state != currentState.UserState {
			return nil
		}
		return c(ctx)
	}
}

func SaveLastMessage(c func(tele.Context) ([]tele.Editable, error)) func(tele.Context) error {
	return func(ctx tele.Context) error {
		f := ctx.Get("fsm").(fsm.FSM)
		ctx.Delete()
		result, err := f.GetState(uint64(ctx.Chat().ID))
		currentState, _ := result.(stateconst.State)
		if err != nil {
			nullState := stateconst.NullState
			currentState.UserState = nullState
			currentState.PreviousMessages = nil
			f.SetState(uint64(ctx.Chat().ID), currentState)
		}
		messages, err := c(ctx)
		if err != nil {
			return err
		}
		currentState.CurrentMessages = messages
		f.SetState(uint64(ctx.Chat().ID), currentState)
		return nil
	}
}

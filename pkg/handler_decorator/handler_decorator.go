package handlerdecorator

import (
	"github.com/avi-gecko/fsm/pkg/fsm"
	"github.com/RG1ee/gobot/pkg/user_state"
	tele "gopkg.in/telebot.v3"
)

func DecoratorHandle(c func(tele.Context) error, state userstate.UserState) func(tele.Context) error {
	return func(ctx tele.Context) error {
		f := ctx.Get("fsm").(fsm.FSM)
		currentState, err := f.GetState(uint64(ctx.Chat().ID))
		if err != nil {
			return nil
		}
		if state != currentState {
			return nil
		}
		return c(ctx)
	}
}

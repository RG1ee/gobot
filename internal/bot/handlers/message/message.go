package message

import (
	"fmt"

	"github.com/RG1ee/gobot/internal/bot/keyboards/reply"
	stateconst "github.com/RG1ee/gobot/internal/bot/state_const"
	"github.com/avi-gecko/fsm/pkg/fsm"
	tele "gopkg.in/telebot.v3"
)

func StartMessageHandler(c tele.Context) error {
	return c.Send("Привет! Я бот для химчистки", reply.StartKeyboard())
}

func WriteNewClothMessageHandler(c tele.Context) error {
	fsm := c.Get("fsm").(fsm.FSM)
	fsm.SetState(uint64(c.Chat().ID), stateconst.StateSendCloth)
	newState, err := fsm.GetState(uint64(c.Chat().ID))
	if err != nil {
		return err
	}

	return c.Send(fmt.Sprint(newState))
}

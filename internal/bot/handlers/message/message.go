package message

import (
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
	fsm.SetState(uint64(c.Chat().ID), stateconst.StateWaitPhoto)

	return c.Send("Отправьте фото")
}

func GetPhotoClothMessageHandler(c tele.Context) error {
	// fsm := c.Get("fsm").(fsm.FSM)
	// fsm.SetState(uint64(c.Chat().ID), stateconst.StateWaitPhoto)

	// TODO: Add to database photoID and caption

	// NOTE: Get photo ID
	// photoId := c.Message().Photo.FileID

	// NOTE: Get caption
	// messageText := c.Message().Caption

	return c.Send("Hello")
}

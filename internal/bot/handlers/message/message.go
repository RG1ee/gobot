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
	fsm.SetState(uint64(c.Chat().ID), stateconst.StateWaitPhoto)

	return c.Send("Отправьте фото с подписью", reply.CancelKeyboard())
}

func CancelHandler(c tele.Context) error {
	fsm := c.Get("fsm").(fsm.FSM)

	_, err := fsm.GetState(uint64(c.Chat().ID))
	if err != nil {
		return c.Send("Отменять нечего :)")
	}

	// TODO: Add handle the error if there is no state
	fsm.ClearState(uint64(c.Chat().ID))

	return c.Send("Отменил отправку в химчистку", reply.StartKeyboard())
}

func GetPhotoClothMessageHandler(c tele.Context) error {
	// TODO: Add to database photoID and caption
	// NOTE: Get photo ID
	// photoId := c.Message().Photo.FileID

	// NOTE: Get caption
	captionText := c.Message().Caption
	if captionText == "" {
		return c.Send("Отправьте фотографию с подписью")
	}

	return c.Send("Сохранил "+fmt.Sprint(captionText), reply.StartKeyboard())
}

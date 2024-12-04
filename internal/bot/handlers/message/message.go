package message

import (
	"fmt"
	"time"

	"github.com/RG1ee/gobot/internal/bot/keyboards/reply"
	stateconst "github.com/RG1ee/gobot/internal/bot/state_const"
	"github.com/RG1ee/gobot/internal/repository"
	"github.com/RG1ee/gobot/pkg/domain"
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
		return c.Send("Отменять нечего :)", reply.StartKeyboard())
	}

	fsm.ClearState(uint64(c.Chat().ID))

	return c.Send("Отменил отправку в химчистку", reply.StartKeyboard())
}

func GetPhotoClothMessageHandler(c tele.Context) error {
	db := c.Get("repository").(repository.Cloth)
	fsm := c.Get("fsm").(fsm.FSM)

	photoId := c.Message().Photo.FileID
	captionText := c.Message().Caption
	if captionText == "" {
		return c.Send("Отправьте фотографию с подписью")
	}

	insertData := domain.Cloth{
		Name:         captionText,
		PhotoId:      photoId,
		IncomingDate: time.Now(),
		Status:       domain.ClothIncoming,
	}
	db.Insert(insertData)
	fsm.ClearState(uint64(c.Chat().ID))
	return c.Send("Вещь "+fmt.Sprint(captionText)+" отправлена", reply.StartKeyboard())
}

func GetListIncomingClothMessageHandler(c tele.Context) error {
	db := c.Get("repository").(repository.Cloth)

	allCloth := db.GetIncoming()
	photo := &tele.Photo{File: tele.File{FileID: allCloth[0].PhotoId}, Caption: allCloth[0].Name}
	_, err := photo.Send(c.Bot(), c.Chat(), &tele.SendOptions{ReplyMarkup: reply.CancelKeyboard()})
	return err
}

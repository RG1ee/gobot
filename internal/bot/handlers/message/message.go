package message

import (
	"fmt"
	"time"

	"github.com/RG1ee/gobot/internal/bot/keyboards/inline"
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

func handleClothList(c tele.Context, getClothFunc func() []domain.Cloth, isOutgoing bool) error {
	page := 0
	pageSize := 1

	allCloth := getClothFunc()
	if len(allCloth) == 0 {
		return c.Send("Нет отправленных вещей")
	}

	paginationKeyboard := inline.GeneratePaginationKeyboard(allCloth, page, pageSize, isOutgoing)

	photo := &tele.Photo{
		File:    tele.File{FileID: allCloth[page].PhotoId},
		Caption: allCloth[page].Name,
	}

	return c.Send(photo, &tele.SendOptions{ReplyMarkup: paginationKeyboard})
}

func GetListIncomingClothMessageHandler(c tele.Context) error {
	db := c.Get("repository").(repository.Cloth)
	return handleClothList(c, db.GetIncoming, false)
}

func GetListOutgoingClothMessageHandler(c tele.Context) error {
	db := c.Get("repository").(repository.Cloth)
	return handleClothList(c, db.GetOutgoing, true)
}

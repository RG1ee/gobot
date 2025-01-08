package message

import (
	// "fmt"
	// "time"
	//
	// "github.com/RG1ee/gobot/internal/bot/keyboards/inline"

	"github.com/RG1ee/gobot/internal/bot/keyboards/reply"
	"github.com/avi-gecko/fsm/pkg/fsm"

	stateconst "github.com/RG1ee/gobot/internal/bot/state_const"
	// "github.com/RG1ee/gobot/internal/repository"
	// "github.com/RG1ee/gobot/pkg/domain"
	// "github.com/avi-gecko/fsm/pkg/fsm"
	tele "gopkg.in/telebot.v3"
)

func StartMessageHandler(c tele.Context) ([]tele.Editable, error) {
	message, err := c.Bot().Send(c.Chat(), "Привет! Я бот для химчистки", reply.StartKeyboard())
	return []tele.Editable{message}, err
}

func WriteNewClothMessageHandler(c tele.Context) ([]tele.Editable, error) {
	fsm := c.Get("fsm").(fsm.FSM)
	result, err := fsm.GetState(uint64(c.Chat().ID))
	if err != nil {
		return nil, err
	}
	currentState, ok := result.(stateconst.State)
	if !ok {
		panic(ok)
	}
	currentState.UserState = stateconst.StateWaitPhoto
	fsm.SetState(uint64(c.Chat().ID), currentState)
	message, err := c.Bot().Send(c.Chat(), "Отправьте фото с подписью", reply.CancelKeyboard())
	return []tele.Editable{message}, err
}

// func CancelHandler(c tele.Context) ([]tele.Editable, error) {
// 	fsm := c.Get("fsm").(fsm.FSM)
// 	_, err := fsm.GetState(uint64(c.Chat().ID))
//
// 	if err != nil {
// 		return c.Send("Отменять нечего :)", reply.StartKeyboard())
// 	}
//
// 	fsm.ClearState(uint64(c.Chat().ID))
//
// 	return c.Send("Отменил отправку в химчистку", reply.StartKeyboard())
// }
//
// func GetPhotoClothMessageHandler(c tele.Context) ([]tele.Editable, error) {
// 	db := c.Get("repository").(repository.Cloth)
// 	fsm := c.Get("fsm").(fsm.FSM)
//
// 	photoId := c.Message().Photo.FileID
// 	captionText := c.Message().Caption
// 	if captionText == "" {
// 		return c.Send("Отправьте фотографию с подписью")
// 	}
//
// 	insertData := domain.Cloth{
// 		Name:         captionText,
// 		PhotoId:      photoId,
// 		IncomingDate: time.Now(),
// 		Status:       domain.ClothIncoming,
// 	}
// 	db.Insert(insertData)
// 	fsm.ClearState(uint64(c.Chat().ID))
// 	return c.Send("Вещь "+fmt.Sprint(captionText)+" отправлена", reply.StartKeyboard())
// }
//
// func handleClothList(c tele.Context, getClothFunc func() []domain.Cloth, isOutgoing bool, uniquePrevBtn string, uniqueNextBtn string) ([]tele.Editable, error) {
// 	page := 0
// 	pageSize := 1
//
// 	allCloth := getClothFunc()
// 	if len(allCloth) == 0 {
// 		return c.Send("Нет отправленных вещей")
// 	}
//
// 	paginationKeyboard := inline.GeneratePaginationKeyboard(allCloth, page, pageSize, isOutgoing, uniquePrevBtn, uniqueNextBtn)
//
// 	photo := &tele.Photo{
// 		File:    tele.File{FileID: allCloth[page].PhotoId},
// 		Caption: allCloth[page].Name,
// 	}
//
// 	return c.Send(photo, &tele.SendOptions{ReplyMarkup: paginationKeyboard})
// }
//
// func GetListIncomingClothMessageHandler(c tele.Context) ([]tele.Editable, error) {
// 	db := c.Get("repository").(repository.Cloth)
// 	return handleClothList(c, db.GetIncoming, false, "incoming_prev_btn", "incoming_next_btn")
// }
//
// func GetListOutgoingClothMessageHandler(c tele.Context) ([]tele.Editable, error) {
// 	db := c.Get("repository").(repository.Cloth)
// 	return handleClothList(c, db.GetOutgoing, true, "prev_btn", "next_btn")
// }

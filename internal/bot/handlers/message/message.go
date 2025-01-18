package message

import (
	// "fmt"
	// "time"
	//
	"github.com/RG1ee/gobot/internal/bot/keyboards/inline"

	"time"

	"github.com/RG1ee/gobot/internal/bot/keyboards/reply"
	"github.com/RG1ee/gobot/internal/repository"
	"github.com/RG1ee/gobot/pkg/domain"
	"github.com/avi-gecko/fsm/pkg/fsm"

	"github.com/RG1ee/gobot/pkg/component_middlewares"
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
	fsm := c.Get("fsm").(fsm.FSM[component_middlewares.State])
	currentState, err := fsm.GetState(uint64(c.Chat().ID))
	if err != nil {
		return nil, err
	}
	currentState.UserState = component_middlewares.StateWaitPhoto
	fsm.SetState(uint64(c.Chat().ID), currentState)
	message, err := c.Bot().Send(c.Chat(), "Отправьте фото с подписью", reply.CancelKeyboard())
	return []tele.Editable{message}, err
}

func CancelHandler(c tele.Context) ([]tele.Editable, error) {
	fsm := c.Get("fsm").(fsm.FSM[component_middlewares.State])
	currentState, err := fsm.GetState(uint64(c.Chat().ID))
	if err != nil {
		panic(err)
	}
	if currentState.UserState == component_middlewares.NullState {
		message, err := c.Bot().Send(c.Chat(), "Главное меню", reply.StartKeyboard())
		return []tele.Editable{message}, err
	}

	fsm.SetState(uint64(c.Chat().ID), currentState)
	message, err := c.Bot().Send(c.Chat(), "Главное меню", reply.StartKeyboard())
	return []tele.Editable{message}, err
}

func GetPhotoClothMessageHandler(c tele.Context) ([]tele.Editable, error) {
	db := c.Get("repository").(repository.Cloth)
	fsm := c.Get("fsm").(fsm.FSM[component_middlewares.State])

	photoId := c.Message().Photo.FileID
	captionText := c.Message().Caption
	if captionText == "" {
		message, err := c.Bot().Send(c.Chat(), "Отправьте фото с подписью", reply.CancelKeyboard())
		return []tele.Editable{message}, err
	}

	insertData := domain.Cloth{
		Name:         captionText,
		PhotoId:      photoId,
		IncomingDate: time.Now(),
		Status:       domain.ClothIncoming,
	}
	db.Insert(insertData)
	currentState, _ := fsm.GetState(uint64(c.Chat().ID))
	currentState.UserState = component_middlewares.NullState
	fsm.SetState(uint64(c.Chat().ID), currentState)
	message, err := c.Bot().Send(c.Chat(), "Вещь добавлена успешно", reply.StartKeyboard())
	return []tele.Editable{message}, err
}

func GetListIncomingClothMessageHandler(c tele.Context) ([]tele.Editable, error) {
	fsm := c.Get("fsm").(fsm.FSM[component_middlewares.State])
	db := c.Get("repository").(repository.Cloth)
	db.ClearRotten()

	allCloth := db.GetIncoming()
	messagesList := []tele.Editable{}
	currentState, _ := fsm.GetState(uint64(c.Chat().ID))
	currentState.UserState = component_middlewares.StateSaveChanges
	fsm.SetState(uint64(c.Chat().ID), currentState)

	if len(allCloth) == 0 {
		message, err := c.Bot().Send(c.Chat(), "Список вещей пуст", reply.StartKeyboard())
		return []tele.Editable{message}, err
	}

	messageListClothes, _ := c.Bot().Send(c.Chat(), "Список вещей", reply.SaveChangesKeyboard())
	for _, cloth := range allCloth {
		photo := &tele.Photo{
			File:    tele.File{FileID: cloth.PhotoId},
			Caption: cloth.Name,
		}
		message, _ := c.Bot().Send(c.Chat(), photo, inline.DeleteKeyboard(int(cloth.ID)))
		messagesList = append(messagesList, message, messageListClothes)
	}
	return messagesList, nil
}

func GetListOutClothMessageHandler(c tele.Context) ([]tele.Editable, error) {
	db := c.Get("repository").(repository.Cloth)
	allCloth := db.GetOutgoing()
	messagesList := []tele.Editable{}
	for _, cloth := range allCloth {
		photo := &tele.Photo{
			File:    tele.File{FileID: cloth.PhotoId},
			Caption: cloth.Name,
		}
		message, _ := c.Bot().Send(c.Chat(), photo, reply.StartKeyboard())
		messagesList = append(messagesList, message)
	}
	return messagesList, nil
}

func SaveChangesHandler(c tele.Context) ([]tele.Editable, error) {
	fsm := c.Get("fsm").(fsm.FSM[component_middlewares.State])
	db := c.Get("repository").(repository.Cloth)

	currentState, _ := fsm.GetState(uint64(c.Chat().ID))

	for _, clothId := range currentState.DeletedIdMessages {
		clouth, _ := db.GetById(clothId)
		db.Out(clouth)
	}

	message, _ := c.Bot().Send(c.Chat(), "Изменения сохранены", reply.StartKeyboard())
	messagesList := append(currentState.CurrentMessages, message)
	return messagesList, nil
}

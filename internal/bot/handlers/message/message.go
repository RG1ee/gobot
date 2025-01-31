package message

import (
	"fmt"

	"github.com/RG1ee/gobot/internal/bot/keyboards/inline"

	"time"

	"github.com/RG1ee/gobot/internal/bot/keyboards/reply"
	"github.com/RG1ee/gobot/internal/repository"
	"github.com/RG1ee/gobot/pkg/domain"
	"github.com/avi-gecko/fsm/pkg/fsm"

	"github.com/RG1ee/gobot/pkg/component_middlewares"
	tele "gopkg.in/telebot.v3"
)

func StartHandler(c tele.Context) ([]tele.Editable, error) {
	message, err := c.Bot().Send(c.Chat(), "Привет! Я бот для того, чтобы отслеживать вещи, которые отправляют в химчистку, на починку или реставрацию.", reply.StartKeyboard())
	return []tele.Editable{message}, err
}

func WriteNewClothHandler(c tele.Context) ([]tele.Editable, error) {
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

func GetPhotoClothHandler(c tele.Context) ([]tele.Editable, error) {
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

func GetListIncomingClothHandler(c tele.Context) ([]tele.Editable, error) {
	fsm := c.Get("fsm").(fsm.FSM[component_middlewares.State])
	db := c.Get("repository").(repository.Cloth)

	allCloth := db.GetIncoming()
	messagesList := []tele.Editable{}
	currentState, _ := fsm.GetState(uint64(c.Chat().ID))
	currentState.UserState = component_middlewares.StateSaveChanges
	fsm.SetState(uint64(c.Chat().ID), currentState)

	if len(allCloth) == 0 {
		message, err := c.Bot().Send(c.Chat(), "Список отправленных вещей пуст", reply.StartKeyboard())
		return []tele.Editable{message}, err
	}

	messageListClothes, _ := c.Bot().Send(c.Chat(), "Список отправленных вещей", reply.SaveChangesKeyboard())
	for _, cloth := range allCloth {
		photo := &tele.Photo{
			File:    tele.File{FileID: cloth.PhotoId},
			Caption: fmt.Sprintf("<b>Название:</b> %s\n<b>Дата и время отправления:</b> %s", cloth.Name, cloth.IncomingDate.Format("02-01-2006 15:04:05")),
		}
		message, _ := c.Bot().Send(c.Chat(), photo, inline.DeleteKeyboard(int(cloth.ID)))
		messagesList = append(messagesList, message, messageListClothes)
	}
	return messagesList, nil
}

func GetListOutgoingClothLastSevenDaysHandler(c tele.Context) ([]tele.Editable, error) {
	db := c.Get("repository").(repository.Cloth)
	allCloth := db.GetOutgoingLastSevenDays()
	if len(allCloth) == 0 {
		message, err := c.Bot().Send(c.Chat(), "Список пришедших вещей за 7 дней пуст", reply.StartKeyboard())
		return []tele.Editable{message}, err
	}
	messagesList := []tele.Editable{}
	for _, cloth := range allCloth {
		photo := &tele.Photo{
			File:    tele.File{FileID: cloth.PhotoId},
			Caption: fmt.Sprintf("<b>Название:</b> %s\n<b>Дата и время отправления:</b> %s\n<b>Дата и время прихода:</b> %s", cloth.Name, cloth.IncomingDate.Format("02-01-2006 15:04:05"), cloth.OutgoingDate.Format("02-01-2006 15:04:05")),
		}
		message, _ := c.Bot().Send(c.Chat(), photo, reply.StartKeyboard())
		messagesList = append(messagesList, message)
	}
	return messagesList, nil
}

func GetListOutgoingClothAllTimeHandler(c tele.Context) ([]tele.Editable, error) {
	db := c.Get("repository").(repository.Cloth)
	allCloth := db.GetOutgoingLastSevenDays()
	if len(allCloth) == 0 {
		message, err := c.Bot().Send(c.Chat(), "Список пришедших вещей за все время пуст", reply.StartKeyboard())
		return []tele.Editable{message}, err
	}
	messagesList := []tele.Editable{}
	for _, cloth := range allCloth {
		photo := &tele.Photo{
			File:    tele.File{FileID: cloth.PhotoId},
			Caption: fmt.Sprintf("<b>Название:</b> %s\n<b>Дата и время отправления:</b> %s\n<b>Дата и время прихода:</b> %s", cloth.Name, cloth.IncomingDate.Format("02-01-2006 15:04:05"), cloth.OutgoingDate.Format("02-01-2006 15:04:05")),
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

package callback

import (
	"strconv"

	"github.com/RG1ee/gobot/internal/bot/keyboards/inline"
	utils_app "github.com/RG1ee/gobot/internal/bot/utils"
	"github.com/RG1ee/gobot/internal/repository"
	"github.com/RG1ee/gobot/pkg/component_middlewares"
	"github.com/RG1ee/gobot/pkg/domain"
	"github.com/avi-gecko/fsm/pkg/fsm"
	tele "gopkg.in/telebot.v3"
)

func handlePagination(c tele.Context, allCloth []domain.Cloth, emptyMessage string, isOutgoing bool, uniquePrevBtn string, uniqueNextBtn string) error {
	if len(allCloth) == 0 {
		c.Delete()
		return c.Send(emptyMessage)
	}

	page, _ := strconv.Atoi(c.Callback().Data)
	if page < 0 || page >= len(allCloth) {
		return c.Send("Неверная страница.")
	}

	pageSize := 1
	paginationKeyboard := inline.GeneratePaginationKeyboard(allCloth, page, pageSize, isOutgoing, uniquePrevBtn, uniqueNextBtn)

	photo := &tele.Photo{
		File:    tele.File{FileID: allCloth[page].PhotoId},
		Caption: allCloth[page].Name,
	}
	return c.Edit(photo, &tele.SendOptions{ReplyMarkup: paginationKeyboard})
}

func HandleIncomingPagination(c tele.Context) error {
	db := c.Get("repository").(repository.Cloth)
	allCloth := db.GetIncoming()
	return handlePagination(c, allCloth, "Нет входящих вещей", false, "incoming_prev_btn", "incoming_next_btn")
}

func HandleOutgoingPagination(c tele.Context) error {
	db := c.Get("repository").(repository.Cloth)
	allCloth := db.GetOutgoing()
	return handlePagination(c, allCloth, "Нет отправленных вещей", true, "prev_btn", "next_btn")
}

func IncomingClothHandle(c tele.Context) error {
	fsm := c.Get("fsm").(fsm.FSM[component_middlewares.State])
	db := c.Get("repository").(repository.Cloth)
	clothId, _ := strconv.Atoi(c.Callback().Data)

	currentState, _ := fsm.GetState(uint64(c.Chat().ID))
	currentState.DeletedIdMessages = append(currentState.DeletedIdMessages, clothId)
	fsm.SetState(uint64(c.Chat().ID), currentState)

	photoId, _ := db.GetById(clothId)
	photo := &tele.Photo{
		File:    tele.File{FileID: photoId.PhotoId},
		Caption: "Вещь удалена",
	}
	_, err := c.Bot().Edit(c.Message(), photo, &tele.SendOptions{ReplyTo: c.Message(), ReplyMarkup: inline.ReturnKeyboard(clothId)})
	return err
}

func CancelIncomingClothHandle(c tele.Context) error {
	fsm := c.Get("fsm").(fsm.FSM[component_middlewares.State])
	db := c.Get("repository").(repository.Cloth)

	clothId, _ := strconv.Atoi(c.Callback().Data)
	currentState, _ := fsm.GetState(uint64(c.Chat().ID))

	indexId := utils_app.FindIndex(currentState.DeletedIdMessages, clothId)
	currentState.DeletedIdMessages = append(currentState.DeletedIdMessages[:indexId], currentState.DeletedIdMessages[indexId+1:]...)
	fsm.SetState(uint64(c.Chat().ID), currentState)

	cloth, _ := db.GetById(clothId)
	photo := &tele.Photo{
		File:    tele.File{FileID: cloth.PhotoId},
		Caption: cloth.Name,
	}
	_, err := c.Bot().Edit(c.Message(), photo, &tele.SendOptions{ReplyTo: c.Message(), ReplyMarkup: inline.DeleteKeyboard(clothId)})
	return err
}
